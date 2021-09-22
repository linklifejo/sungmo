# 일봉 데이터 다운
def req_day_data(self, code, repeat=60):
    self.repeat = repeat
    self.dfs = []
    date = str(datetime.date.today()).replace('-', '')
    self.setInputValue("종목코드", code)
    self.setInputValue("기준일자", date)
    self.setInputValue("수정주가구분", "1")
    self.commRqData("일봉데이터요청", "opt10081", 0, "0101")
    return self.dfs

def req_minute_data(self, code, period, repeat=1):
    self.dfs = []
    self.setInputValue("종목코드", code)
    self.setInputValue("틱범위", period)
    self.setInputValue("수정주가구분", "0")
    self.commRqData("분봉데이터요청", "opt10080", 0, "0101")
    for i in range(repeat):
        time.sleep(TR_REQ_TIME_INTERVAL)
        self.setInputValue("종목코드", code)
        self.setInputValue("틱범위", period)
        self.setInputValue("수정주가구분", "0")
        self.commRqData("분봉데이터요청", "opt10080", 2, "0101")
    return self.dfs

def opt10080(self, rqname, trcode):
    dataList = []
    for i in range(self.repeat):
        try:
            day = self.to_datetime(self.commGetData(trcode, "", rqname, i, "체결시간"))
        except:
            df = DataFrame(data=dataList, columns=['Date', 'Open', 'High', 'Low', 'Close', 'Volume'])
            self.dfs.append(df)
            return

        open = self.commGetData(trcode, "", rqname, i, "시가")
        low = self.commGetData(trcode, "", rqname, i, "저가")
        high = self.commGetData(trcode, "", rqname, i, "고가")
        close = self.commGetData(trcode, "", rqname, i, "현재가")
        volume = self.commGetData(trcode, "", rqname, i, "거래량")
        if open[0] == '-':
            open = open[1:]
        if high[0] == '-':
            high = high[1:]
        if low[0] == '-':
            low = low[1:]
        if close[0] == '-':
            close = close[1:]

        dataList.append([day, int(open), int(high), int(low), int(close), int(volume)])
    df = DataFrame(data=dataList, columns=['Date', 'Open', 'High', 'Low', 'Close', 'Vloume'])
    self.dfs.append(df)
C:\Users\chosm\AppData\Local\Programs\Python\Python37-32