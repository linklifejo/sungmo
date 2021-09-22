import sys
from PyQt5.QtWidgets import *
from PyQt5.QAxContainer import *
from PyQt5.QtCore import *
import time
import pandas as pd
import sqlite3

TR_REQ_TIME_INTERVAL = 0.2

class Kiwoom(QAxWidget):
    def __init__(self):
        super().__init__()
        self._create_kiwoom_instance()
        self._set_signal_slots()

    def _create_kiwoom_instance(self):
        self.setControl("KHOPENAPI.KHOpenAPICtrl.1")

    def _set_signal_slots(self):
        self.OnEventConnect.connect(self._event_connect)

    def comm_connect(self):
        self.dynamicCall("CommConnect()")
        self.login_event_loop = QEventLoop()
        self.login_event_loop.exec_()

    def _event_connect(self, err_code):
        if err_code == 0:
            print("connected")
        else:
            print("disconnected")

        self.login_event_loop.exit()

    def getCodesOfMarket(self, market):
        code_list = self.dynamicCall("GetCodeListByMarket(QString)", market)
        code_list = code_list.split(';')
        return code_list[:-1]

    def getNameOfCode(self, code):
        code_name = self.dynamicCall("GetMasterCodeName(QString)", code)
        return code_name

    def setInputValue(self, id, value):
        self.dynamicCall("SetInputValue(QString, QString)", id, value)

    def commRqData(self, rqname, trcode, next, screenNo):
        self.dynamicCall("CommRqData(QString, QString, int, QString", rqname, trcode, next, screenNo)
        self.trEeventLoop = QEventLoop()
        self.trEventLoop.exec_()

    def commGetData(self, code, realType, fieldName, index, itemName):
        ret = self.dynamicCall("CommGetData(QString, QString, QString, int, QString", code,
                               realType, fieldName, index, itemName)
        return ret.strip()

    def getRepeatCnt(self, trcode, rqname):
        ret = self.dynamicCall("GetRepeatCnt(QString, QString)", trcode, rqname)
        return ret

    def receiveTrData(self, screenNO, rqname, trcode, recordName, next, unused1, unused2, unused3, unused4):
        if next == '2':
            self.remainedData = True
        else:
            self.remainedData = False

        if rqname == "opt10081_req":
            self.opt10081(rqname, trcode)

        try:
            self.tr_event_loop.exit()
        except AttributeError:
            pass

    # 일봉 데이터 다운
    def req_day_data(self, code, repeat=60):
        # opt10081 TR 요청
        self.dfs = []
        self.repeat = repeat
        date = str(datetime.date.today()).replace('-', '')
        self.setInputValue("종목코드", code)
        self.setInputValue("기준일자", date)
        self.setInputValue("수정주가구분", "1")
        self.commRqData("일봉데이터요청", "opt10081", 0, "0101")
        while self.cnt == True:
            time.sleep(TR_REQ_TIME_INTERVAL)
            self.setInputValue("종목코드", code)
            self.setInputValue("기준일자", date)
            self.setInputValue("수정주가구분", "1")
            self.commRqData("일봉데이터요청", "opt10081", 2, "0101")
        return self.dfs

    def opt10081(self, rqname, trcode):
        cnt = self.getRepeatCnt(trcode, rqname)

        for i in range(cnt):
            date = self._comm_get_data(trcode, "", rqname, i, "일자")
            open = self._comm_get_data(trcode, "", rqname, i, "시가")
            high = self._comm_get_data(trcode, "", rqname, i, "고가")
            low = self._comm_get_data(trcode, "", rqname, i, "저가")
            close = self._comm_get_data(trcode, "", rqname, i, "현재가")
            volume = self._comm_get_data(trcode, "", rqname, i, "거래량")
            print(date, open, high, low, close, volume)


if __name__ == "__main__":
    app = QApplication(sys.argv)
    kiwoom = Kiwoom()
    kiwoom.comm_connect()
    # codeList = kiwoom.getCodesOfMarket('10')
    # for code in codeList:
    #     print(code, end=" ")
    # print(kiwoom.getNameOfCode("000660"))
    # opt10081 TR 요청
    kiwoom.set_input_value("종목코드", "039490")
    kiwoom.set_input_value("기준일자", "20170224")
    kiwoom.set_input_value("수정주가구분", 1)
    kiwoom.comm_rq_data("opt10081_req", "opt10081", 0, "0101")

    while kiwoom.remained_data == True:
        time.sleep(TR_REQ_TIME_INTERVAL)
        kiwoom.set_input_value("종목코드", "039490")
        kiwoom.set_input_value("기준일자", "20170224")
        kiwoom.set_input_value("수정주가구분", 1)
        kiwoom.comm_rq_data("opt10081_req", "opt10081", 2, "0101")
    df = pd.DataFrame(kiwoom.ohlcv, columns=['open', 'high', 'low', 'close', 'volume'], index=kiwoom.ohlcv['date'])
    con = sqlite3.connect("c:/stock/database/daycandle.db")
    df.to_sql('039490', con, if_exists='replace')
