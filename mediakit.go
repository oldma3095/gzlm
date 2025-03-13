package mediakit

import (
	"fmt"
	"gzlm/request"
	"gzlm/response"
	"net/url"
)

const (
	__DefaultVhost__ = "__defaultVhost__"
)

type MediaKit struct {
	Addr   string
	Secret string
}

func NewMediaKit(addr string, secret string) *MediaKit {
	return &MediaKit{Addr: addr, Secret: secret}
}

func (m *MediaKit) checkInit() (err error) {
	if m.Addr == "" {
		err = fmt.Errorf("连接地址为空")
		return
	}
	if m.Secret == "" {
		err = fmt.Errorf("secret为空")
		return
	}
	return
}

func (m *MediaKit) AddFFMpegSource(info request.AddFFMpegSource) (res response.AddFFMpegSource, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.SrcUrl == "" || info.DstUrl == "" {
		err = fmt.Errorf("参数错误, src_url:%s, dst_url:%s", info.SrcUrl, info.DstUrl)
		return
	}

	if info.EnableHls > 0 {
		info.EnableHls = 1
	}

	if info.EnableMp4 > 0 {
		info.EnableMp4 = 1
	}

	uri := "/index/api/addFFmpegSource"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) AddStreamProxy(info request.AddStreamProxy) (res response.AddStreamProxy, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.APP == "" || info.Stream == "" || info.URL == "" {
		err = fmt.Errorf("参数错误, app:%s, stream:%s, url:%s", info.APP, info.Stream, info.URL)
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.RtpType > 2 {
		info.RtpType = 0
	}
	if info.EnableHls > 0 {
		info.EnableHls = 1
	}
	if info.EnableMp4 > 0 {
		info.EnableMp4 = 1
	}
	if info.EnableRtsp > 0 {
		info.EnableRtsp = 1
	}
	if info.EnableRtmp > 0 {
		info.EnableRtmp = 1
	}
	if info.EnableTs > 0 {
		info.EnableTs = 1
	}
	if info.EnableFmp4 > 0 {
		info.EnableFmp4 = 1
	}
	if info.EnableAudio > 0 {
		info.EnableAudio = 1
	}
	if info.AddMuteAudio > 0 {
		info.AddMuteAudio = 1
	}
	if info.ModifyStamp > 0 {
		info.ModifyStamp = 1
	}

	uri := "/index/api/addStreamProxy"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) AddStreamPusherProxy(info request.AddStreamPusherProxy) (res response.AddStreamPusherProxy, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Schema == "" || info.APP == "" || info.Stream == "" || info.DstUrl == "" {
		err = fmt.Errorf("参数错误, schema:%s, app:%s, stream:%s, url:%s", info.Schema, info.APP, info.Stream, info.DstUrl)
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.RtpType > 2 {
		info.RtpType = 0
	}

	uri := "/index/api/addStreamPusherProxy"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) CloseRtpServer(info request.CloseRtpServer) (res response.CloseRtpServer, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.StreamId == "" {
		err = fmt.Errorf("参数错误, stream_id:%s", info.StreamId)
		return
	}

	uri := "/index/api/closeRtpServer"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) CloseStreams(info request.CloseStreams) (res response.CloseStreams, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Force > 0 {
		info.Force = 1
	}

	uri := "/index/api/close_streams"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) DelFFMpegSource(info request.DelFFMpegSource) (res response.DelFFMpegSource, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Key == "" {
		err = fmt.Errorf("key不能为空")
		return
	}

	uri := "/index/api/delFFmpegSource"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) DelStreamProxy(info request.DelStreamProxy) (res response.DelStreamProxy, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Key == "" {
		err = fmt.Errorf("key不能为空")
		return
	}

	uri := "/index/api/delStreamProxy"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) DelStreamPusherProxy(info request.DelStreamPusherProxy) (res response.DelStreamPusherProxy, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/delStreamPusherProxy"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetAllSession(info request.GetAllSession) (res response.GetAllSession, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/getAllSession"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetApiList(info request.GetApiList) (res response.GetApiList, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/getApiList"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetMediaList(info request.GetMediaList) (res response.GetMediaList, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/getMediaList"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetMp4RecordFile(info request.GetMp4RecordFile) (res response.GetMp4RecordFile, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.APP == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s", info.Vhost, info.APP, info.Stream)
		return
	}

	uri := "/index/api/getMp4RecordFile"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetRtpInfo(info request.GetRtpInfo) (res response.GetRtpInfo, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.StreamId == "" {
		err = fmt.Errorf("参数错误, stream_id:%s", info.StreamId)
		return
	}

	uri := "/index/api/getRtpInfo"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetServerConfig(info request.GetServerConfig) (res response.GetServerConfig, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/getServerConfig"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetSnap(info request.GetSnap) (res response.GetSnap, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.URL == "" || info.TimeoutSec == 0 || info.ExpireSec == 0 {
		err = fmt.Errorf("参数不能为空或者0")
		return
	}

	uri := "/index/api/getSnap"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetStatistic(info request.GetStatistic) (res response.GetStatistic, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/getStatistic"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetThreadsLoad(info request.GetThreadsLoad) (res response.GetThreadsLoad, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/getThreadsLoad"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetWorkThreadsLoad(info request.GetWorkThreadsLoad) (res response.GetWorkThreadsLoad, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/getWorkThreadsLoad"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) IsMediaOnline(info request.IsMediaOnline) (res response.IsMediaOnline, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Schema == "" || info.Vhost == "" || info.APP == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, schema:%s, vhost:%s, app:%s, stream:%s", info.Schema, info.Vhost, info.APP, info.Stream)
		return
	}

	uri := "/index/api/isMediaOnline"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) KickSessions(info request.KickSessions) (res response.KickSessions, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/kick_sessions"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) ListRtpServer(info request.ListRtpServer) (res response.ListRtpServer, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/listRtpServer"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) OpenRtpServer(info request.OpenRtpServer) (res response.OpenRtpServer, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.StreamId == "" {
		err = fmt.Errorf("stream_id不能为空")
		return
	}

	if info.Port > 65535 {
		info.Port = 0
	}
	if info.TcpMode > 2 {
		info.TcpMode = 1
	}
	if info.ReUsePort > 1 {
		info.ReUsePort = 0
	}
	if info.OnlyAudio > 1 {
		info.OnlyAudio = 0
	}

	uri := "/index/api/openRtpServer"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) RestartServer(info request.RestartServer) (res response.RestartServer, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/restartServer"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) SetServerConfig(info map[string]string) (res response.SetServerConfig, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}
	uri := "/index/api/setServerConfig"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	if info == nil {
		info = make(map[string]string)
	}
	if _, ok := info["secret"]; !ok && m.Secret != "" {
		info["secret"] = m.Secret
	}
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) StartRecord(info request.StartRecord) (res response.StartRecord, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.APP == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s", info.Vhost, info.APP, info.Stream)
		return
	}

	if info.Type > 1 {
		info.Type = 1
	}

	uri := "/index/api/startRecord"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) StartSendRtp(info request.StartSendRtp) (res response.StartSendRtp, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.APP == "" || info.Stream == "" || info.SSRC == 0 || info.DstUrl == "" ||
		info.DstPort == 0 || info.DstPort > 65535 {
		err = fmt.Errorf(
			"参数错误, vhost:%s, app:%s, stream:%s, ssrc:%d, dst_url:%s, dst_port:%d, src_port:%d",
			info.Vhost, info.APP, info.Stream, info.SSRC, info.DstUrl, info.DstPort, info.DstPort)
		return
	}
	if info.SrcPort > 65535 {
		err = fmt.Errorf("参数错误, src_port:%d", info.DstPort)
		return
	}

	if info.IsUdp > 0 {
		info.IsUdp = 1
	}
	if info.FromMp4 > 0 {
		info.FromMp4 = 1
	}
	if info.UsePs > 0 {
		info.UsePs = 1
	}
	if info.OnlyAudio > 0 {
		info.OnlyAudio = 1
	}
	if info.UdpRtcpTimeout > 0 {
		info.UdpRtcpTimeout = 1
	}

	uri := "/index/api/startSendRtp"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) StopRecord(info request.StopRecord) (res response.StopRecord, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.APP == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s", info.Vhost, info.APP, info.Stream)
		return
	}
	if info.Type > 1 {
		info.Type = 1
	}

	uri := "/index/api/stopRecord"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) StopSendRtp(info request.StopSendRtp) (res response.StopSendRtp, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.APP == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s", info.Vhost, info.APP, info.Stream)
		return
	}

	uri := "/index/api/stopSendRtp"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetProxyInfo(info request.GetProxyInfo) (res response.GetProxyInfo, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Key == "" {
		err = fmt.Errorf("参数不能为空")
		return
	}

	uri := "/index/api/getProxyInfo"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) GetProxyPusherInfo(info request.GetProxyPusherInfo) (res response.GetProxyPusherInfo, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Key == "" {
		err = fmt.Errorf("参数不能为空")
		return
	}

	uri := "/index/api/getProxyPusherInfo"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}

func (m *MediaKit) SetupTranscode(info request.SetupTranscode) (res response.SetupTranscode, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.APP == "" || info.Stream == "" || info.Name == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s, name:%s", info.Vhost, info.APP, info.Stream, info.Name)
		return
	}
	postData := make(map[string]interface{})
	postData["secret"] = info.Secret
	postData["vhost"] = info.Vhost
	postData["app"] = info.APP
	postData["stream"] = info.Stream
	postData["name"] = info.Name

	if info.Add < 0 || info.Add > 1 {
		err = fmt.Errorf("参数错误, add:%d", info.Add)
		return
	}
	postData["add"] = info.Add
	postData["force"] = true

	postData["video_codec"] = info.VideoCodec
	postData["video_bitrate"] = info.VideoBitrate
	if info.VideoScale > 0 && (info.VideoScale < 0.1 || info.VideoScale > 10) {
		err = fmt.Errorf("参数错误 0.1~10, video_scale:%f", info.VideoScale)
		return
	}
	postData["video_scale"] = info.VideoScale
	postData["audio_codec"] = info.AudioCodec
	postData["audio_bitrate"] = info.AudioBitrate
	postData["audio_samplerate"] = info.AudioSamplerate
	postData["hw_decoder"] = false
	postData["hw_encoder"] = false
	postData["decoder_list"] = "libx264"
	postData["encoder_list"] = "libx264"

	uri := "/index/api/setupTranscode"
	u, err := url.JoinPath(m.Addr, uri)
	if err != nil {
		return
	}

	req := Request()
	req.SetResult(&res)
	req.SetQueryParam("secret", m.Secret)
	req.SetBody(info)
	_, err = req.Post(u)
	if err != nil {
		return
	}
	return
}
