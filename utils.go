/*
 * Copyright 2022 douyincloud
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dyclog

import (
	"net"
	"net/http"
	"runtime"
	"strings"
)

func GetCallerLocation(caller *runtime.Frame) (string, int) {
	file := caller.File
	line := caller.Line

	lastSlash := strings.LastIndex(file, "/")
	if lastSlash == -1 || lastSlash == len(file)-1 {
		return "", line
	}
	return file[lastSlash+1:], line
}

func GetRemoteIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetCaller(depth int) *runtime.Frame {
	pcs := make([]uintptr, 10)
	_ = runtime.Callers(depth+1, pcs)
	frames := runtime.CallersFrames(pcs)
	f, _ := frames.Next()
	return &f
}
