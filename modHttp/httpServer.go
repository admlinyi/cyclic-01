package modHttp

type CHttpServer struct {

}

var g_singleHttpServer *CHttpServer = &CHttpServer{}

func getSingleHttpServer() *CHttpServer {
	return g_singleHttpServer
}

