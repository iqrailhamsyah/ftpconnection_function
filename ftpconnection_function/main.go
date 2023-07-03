package main

import (
	"github.com/webguerilla/ftps"
)

func main() {
	koneksiftp := NewFtps()
	kredensial := &credentials{
		host:     "fsrv.bri.co.id",
		port:     50021,
		username: "H2H_JSRFaspay_USR2023",
		password: "w@4rTCaBmpZL",
	}
	direktori := &directorystring{
		localfiledirectory: "C:/Users/User/Downloads/",
		ftpfiledirectory:   "/faspay/settlement_xlsx/Processed/",
		filename:           "JUNIOSMART-2022-10-04-settlement.xlsx",
	}
	koneksiftp.DownloadExcel(kredensial, direktori)
}

func NewFtps() FtpsconnectionFunction {
	//membuat objek ftps
	ftps := new(ftps.FTPS)
	return ftpsconnectionfunction{
		libs: ftps,
	}
}

type FtpsconnectionFunction interface {
	DownloadExcel(cred *credentials, dir *directorystring)
}

type ftpsconnectionfunction struct {
	libs *ftps.FTPS
}

type credentials struct {
	host     string
	port     int
	username string
	password string
}

type directorystring struct {
	localfiledirectory string
	ftpfiledirectory   string
	filename           string
}

func (f ftpsconnectionfunction) DownloadExcel(cred *credentials, dir *directorystring) {

	f.libs.TLSConfig.InsecureSkipVerify = true // often necessary in shared hosting environments
	f.libs.Debug = true

	//setting connect ke host
	err := f.libs.Connect(cred.host, cred.port)
	if err != nil {
		panic(err)
	}

	//menutup koneksi FTP
	defer f.libs.Quit()

	//setting credential username & password
	err = f.libs.Login(cred.username, cred.password)
	if err != nil {
		panic(err)
	}

	//proses download file dari remote FTP server ke local directory
	f.libs.RetrieveFile(dir.ftpfiledirectory+dir.filename, dir.localfiledirectory+dir.filename)

}
