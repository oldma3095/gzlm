package gzlm

import (
	"fmt"
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

func (m *MediaKit) AddFFMpegSource(info AddFFMpegSourceReq) (res AddFFMpegSourceRes, err error) {
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

func (m *MediaKit) AddStreamProxy(info AddStreamProxyReq) (res AddStreamProxyRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.App == "" || info.Stream == "" || info.Url == "" {
		err = fmt.Errorf("参数错误, app:%s, stream:%s, url:%s", info.App, info.Stream, info.Url)
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

func (m *MediaKit) AddStreamPusherProxy(info AddStreamPusherProxyReq) (res AddStreamPusherProxyRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Schema == "" || info.App == "" || info.Stream == "" || info.DstUrl == "" {
		err = fmt.Errorf("参数错误, schema:%s, app:%s, stream:%s, url:%s", info.Schema, info.App, info.Stream, info.DstUrl)
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

func (m *MediaKit) CloseRtpServer(info CloseRtpServerReq) (res CloseRtpServerRes, err error) {
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

func (m *MediaKit) CloseStreams(info CloseStreamsReq) (res CloseStreamsRes, err error) {
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

func (m *MediaKit) DelFFMpegSource(info DelFFMpegSourceReq) (res DelFFMpegSourceRes, err error) {
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

func (m *MediaKit) DelStreamProxy(info DelStreamProxyReq) (res DelStreamProxyRes, err error) {
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

func (m *MediaKit) DelStreamPusherProxy(info DelStreamPusherProxyReq) (res DelStreamPusherProxyRes, err error) {
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

func (m *MediaKit) GetAllSession(info GetAllSessionReq) (res GetAllSessionRes, err error) {
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

func (m *MediaKit) GetApiList(info GetApiListReq) (res GetApiListRes, err error) {
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

func (m *MediaKit) GetMediaList(info GetMediaListReq) (res GetMediaListRes, err error) {
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

func (m *MediaKit) GetMp4RecordFile(info GetMp4RecordFileReq) (res GetMp4RecordFileRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.App == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s", info.Vhost, info.App, info.Stream)
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

func (m *MediaKit) GetRtpInfo(info GetRtpInfoReq) (res GetRtpInfoRes, err error) {
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

func (m *MediaKit) GetServerConfig(info GetServerConfigReq) (res GetServerConfigRes, err error) {
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

func (m *MediaKit) GetSnap(info GetSnapReq) (res GetSnapRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Url == "" || info.TimeoutSec == 0 || info.ExpireSec == 0 {
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

func (m *MediaKit) GetStatistic(info GetStatisticReq) (res GetStatisticRes, err error) {
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

func (m *MediaKit) GetThreadsLoad(info GetThreadsLoadReq) (res GetThreadsLoadRes, err error) {
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

func (m *MediaKit) GetWorkThreadsLoad(info GetWorkThreadsLoadReq) (res GetWorkThreadsLoadRes, err error) {
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

func (m *MediaKit) IsMediaOnline(info IsMediaOnlineReq) (res IsMediaOnlineRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Schema == "" || info.Vhost == "" || info.App == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, schema:%s, vhost:%s, app:%s, stream:%s", info.Schema, info.Vhost, info.App, info.Stream)
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

func (m *MediaKit) KickSessions(info KickSessionsReq) (res KickSessionsRes, err error) {
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

func (m *MediaKit) ListRtpServer(info ListRtpServerReq) (res ListRtpServerRes, err error) {
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

func (m *MediaKit) OpenRtpServer(info OpenRtpServerReq) (res OpenRtpServerRes, err error) {
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

func (m *MediaKit) RestartServer(info RestartServerReq) (res RestartServerRes, err error) {
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

func (m *MediaKit) SetServerConfig(info map[string]string) (res SetServerConfigRes, err error) {
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

func (m *MediaKit) StartRecord(info StartRecordReq) (res StartRecordRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.App == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s", info.Vhost, info.App, info.Stream)
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

func (m *MediaKit) StartSendRtp(info StartSendRtpReq) (res StartSendRtpRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.App == "" || info.Stream == "" || info.Ssrc == 0 || info.DstUrl == "" ||
		info.DstPort == 0 || info.DstPort > 65535 {
		err = fmt.Errorf(
			"参数错误, vhost:%s, app:%s, stream:%s, ssrc:%d, dst_url:%s, dst_port:%d, src_port:%d",
			info.Vhost, info.App, info.Stream, info.Ssrc, info.DstUrl, info.DstPort, info.DstPort)
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

func (m *MediaKit) StopRecord(info StopRecordReq) (res StopRecordRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.App == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s", info.Vhost, info.App, info.Stream)
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

func (m *MediaKit) StopSendRtp(info StopSendRtpReq) (res StopSendRtpRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.App == "" || info.Stream == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s", info.Vhost, info.App, info.Stream)
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

func (m *MediaKit) GetProxyInfo(info GetProxyInfoReq) (res GetProxyInfoRes, err error) {
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

func (m *MediaKit) GetProxyPusherInfo(info GetProxyPusherInfoReq) (res GetProxyPusherInfoRes, err error) {
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

func (m *MediaKit) SetupTranscode(info SetupTranscodeReq) (res SetupTranscodeRes, err error) {
	err = m.checkInit()
	if err != nil {
		return
	}

	if info.Vhost == "" {
		info.Vhost = __DefaultVhost__
	}

	if info.Vhost == "" || info.App == "" || info.Stream == "" || info.Name == "" {
		err = fmt.Errorf("参数错误, vhost:%s, app:%s, stream:%s, name:%s", info.Vhost, info.App, info.Stream, info.Name)
		return
	}
	postData := make(map[string]interface{})
	postData["secret"] = info.Secret
	postData["vhost"] = info.Vhost
	postData["app"] = info.App
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
