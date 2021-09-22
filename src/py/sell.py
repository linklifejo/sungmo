from pytrader import MyWindow
import pandas as pd
import datetime
import json
from get_data import get_data
import sqlite3
import os
import time
from PyQt5.QtCore import QThread, pyqtSignal, QEventLoop
import pause
from order import Order
from dateutil.relativedelta import relativedelta


class Sell(QThread):
    tx_signal = pyqtSignal(str)
    rx_signal = pyqtSignal(str)

    def __init__(self, user_param, cond):
        super(Sell, self).__init__()
        self.user_param = user_param
        self.cond = cond
        self.rx_signal.connect(self.exec_order)

    def __set_time_format(self, date, time):
        return datetime.datetime.strptime(date + ' ' + time, "%Y-%m-%d %H:%M:%S")

    def send(self, command):
        self.tx_signal.emit(command)
        self.orderLoop = QEventLoop()
        self.orderLoop.exec_()

    def done(self):
        if self.orderLoop is not None:
            self.orderLoop.exit()

    def exec_order(self, command):
        if ',' in command:
            args = command.split(',')
            eval('self.' + args[0])(*args[1:])
        else:
            eval('self.' + command)()

    def monitor(self, close_time):
        gd = get_data(self.user_param['path']['root'])
        con = sqlite3.connect(os.path.join(self.user_param['path']['root'], 'database', 'order.db'))
        df = pd.read_sql("SELECT * FROM 'hold_list'", con)
        today = str(datetime.date.today())
        open_time = datetime.datetime.strptime(today + ' ' + self.user_param['market-time']['open'], "%Y-%m-%d %H:%M:%S")
        while open_time <= datetime.datetime.now() < close_time:
            for idx in df.index:
                code = df.loc[idx, 'code']
                tid =  df.loc[idx, 'tid']
                sell_price = df.loc[idx, 'sell_price']
                stop_price = df.loc[idx, 'stop_price']
                due_date = df.loc[idx, 'due_date']
                due_date = datetime.datetime.strptime(due_date, '%Y-%m-%d %H:%M:%S')
                now = datetime.datetime.now()
                trade_price = gd.get_current_price(code)
                if trade_price is not None:
                    if (sell_price - trade_price) / trade_price < 0.01:
                        command = ';'.join([self.__class__.__name__, 'Buy', 'sell_order,{},시장가'.format(tid)])
                        print("익절 시장가 매도처리~~")
                        self.send(command)

                    if trade_price <= stop_price:
                        command = ';'.join([self.__class__.__name__, 'Buy', 'sell_order,{},시장가'.format(tid)])
                        print("시장가 매도처리~~")
                        self.send(command)
                    if due_date <= now:
                        command = ';'.join([self.__class__.__name__, 'Buy', 'sell_order,{},시장가'.format(tid)])
                        self.send(command)
                    time.sleep(0.5)
            time.sleep(2)


    def run(self):
        today = str(datetime.date.today())
        open_time = self.__set_time_format(today, self.user_param['market-time']['open'])
        close_time = self.__set_time_format(today, self.user_param['market-time']['close'])
        dt_time = self.__set_time_format(today, self.user_param['market-time']['daytrade'])
        swing_time = self.__set_time_format(today, self.user_param['market-time']['swing'])
        now = datetime.datetime.now()

        interval = 3
        if open_time <= now < close_time:
            now = Order.roundTime(now, interval)
            print("wait until {}".format(now))
            self.monitor(now)
            # day trading
            while datetime.datetime.now() < close_time - relativedelta(minutes=5):
                # if datetime.datetime.now() < dt_time:
                #     print("wait until {}".format(dt_time))
                #     self.monitor(dt_time)
                command = ';'.join([self.__class__.__name__, 'Buy', 'daytrade'])
                self.send(command)

                now += relativedelta(seconds=interval)
                print("wait until {}".format(now))
                self.monitor(now)

            print("wait until {}".format(swing_time))
            self.monitor(swing_time)
            command = ';'.join([self.__class__.__name__, 'Buy', 'swing'])
            self.send(command)
        else:
            print("market is closed")
