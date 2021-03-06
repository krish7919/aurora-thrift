// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"api"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Response createJob(JobConfiguration description, Lock lock)")
	fmt.Fprintln(os.Stderr, "  Response scheduleCronJob(JobConfiguration description, Lock lock)")
	fmt.Fprintln(os.Stderr, "  Response descheduleCronJob(JobKey job, Lock lock)")
	fmt.Fprintln(os.Stderr, "  Response startCronJob(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response restartShards(JobKey job,  shardIds, Lock lock)")
	fmt.Fprintln(os.Stderr, "  Response killTasks(TaskQuery query, Lock lock, JobKey job,  instances)")
	fmt.Fprintln(os.Stderr, "  Response addInstances(AddInstancesConfig config, Lock lock, InstanceKey key, i32 count)")
	fmt.Fprintln(os.Stderr, "  Response acquireLock(LockKey lockKey)")
	fmt.Fprintln(os.Stderr, "  Response releaseLock(Lock lock, LockValidation validation)")
	fmt.Fprintln(os.Stderr, "  Response replaceCronTemplate(JobConfiguration config, Lock lock)")
	fmt.Fprintln(os.Stderr, "  Response startJobUpdate(JobUpdateRequest request, string message)")
	fmt.Fprintln(os.Stderr, "  Response pauseJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response resumeJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response abortJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response pulseJobUpdate(JobUpdateKey key)")
	fmt.Fprintln(os.Stderr, "  Response getRoleSummary()")
	fmt.Fprintln(os.Stderr, "  Response getJobSummary(string role)")
	fmt.Fprintln(os.Stderr, "  Response getTasksStatus(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getTasksWithoutConfigs(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getPendingReason(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getConfigSummary(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response getJobs(string ownerRole)")
	fmt.Fprintln(os.Stderr, "  Response getQuota(string ownerRole)")
	fmt.Fprintln(os.Stderr, "  Response populateJobConfig(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response getLocks()")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateSummaries(JobUpdateQuery jobUpdateQuery)")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateDetails(JobUpdateKey key)")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateDiff(JobUpdateRequest request)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := api.NewAuroraSchedulerManagerClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "createJob":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "CreateJob requires 2 args")
			flag.Usage()
		}
		arg158 := flag.Arg(1)
		mbTrans159 := thrift.NewTMemoryBufferLen(len(arg158))
		defer mbTrans159.Close()
		_, err160 := mbTrans159.WriteString(arg158)
		if err160 != nil {
			Usage()
			return
		}
		factory161 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt162 := factory161.GetProtocol(mbTrans159)
		argvalue0 := api.NewJobConfiguration()
		err163 := argvalue0.Read(jsProt162)
		if err163 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg164 := flag.Arg(2)
		mbTrans165 := thrift.NewTMemoryBufferLen(len(arg164))
		defer mbTrans165.Close()
		_, err166 := mbTrans165.WriteString(arg164)
		if err166 != nil {
			Usage()
			return
		}
		factory167 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt168 := factory167.GetProtocol(mbTrans165)
		argvalue1 := api.NewLock()
		err169 := argvalue1.Read(jsProt168)
		if err169 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.CreateJob(value0, value1))
		fmt.Print("\n")
		break
	case "scheduleCronJob":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ScheduleCronJob requires 2 args")
			flag.Usage()
		}
		arg170 := flag.Arg(1)
		mbTrans171 := thrift.NewTMemoryBufferLen(len(arg170))
		defer mbTrans171.Close()
		_, err172 := mbTrans171.WriteString(arg170)
		if err172 != nil {
			Usage()
			return
		}
		factory173 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt174 := factory173.GetProtocol(mbTrans171)
		argvalue0 := api.NewJobConfiguration()
		err175 := argvalue0.Read(jsProt174)
		if err175 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg176 := flag.Arg(2)
		mbTrans177 := thrift.NewTMemoryBufferLen(len(arg176))
		defer mbTrans177.Close()
		_, err178 := mbTrans177.WriteString(arg176)
		if err178 != nil {
			Usage()
			return
		}
		factory179 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt180 := factory179.GetProtocol(mbTrans177)
		argvalue1 := api.NewLock()
		err181 := argvalue1.Read(jsProt180)
		if err181 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.ScheduleCronJob(value0, value1))
		fmt.Print("\n")
		break
	case "descheduleCronJob":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DescheduleCronJob requires 2 args")
			flag.Usage()
		}
		arg182 := flag.Arg(1)
		mbTrans183 := thrift.NewTMemoryBufferLen(len(arg182))
		defer mbTrans183.Close()
		_, err184 := mbTrans183.WriteString(arg182)
		if err184 != nil {
			Usage()
			return
		}
		factory185 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt186 := factory185.GetProtocol(mbTrans183)
		argvalue0 := api.NewJobKey()
		err187 := argvalue0.Read(jsProt186)
		if err187 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg188 := flag.Arg(2)
		mbTrans189 := thrift.NewTMemoryBufferLen(len(arg188))
		defer mbTrans189.Close()
		_, err190 := mbTrans189.WriteString(arg188)
		if err190 != nil {
			Usage()
			return
		}
		factory191 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt192 := factory191.GetProtocol(mbTrans189)
		argvalue1 := api.NewLock()
		err193 := argvalue1.Read(jsProt192)
		if err193 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.DescheduleCronJob(value0, value1))
		fmt.Print("\n")
		break
	case "startCronJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "StartCronJob requires 1 args")
			flag.Usage()
		}
		arg194 := flag.Arg(1)
		mbTrans195 := thrift.NewTMemoryBufferLen(len(arg194))
		defer mbTrans195.Close()
		_, err196 := mbTrans195.WriteString(arg194)
		if err196 != nil {
			Usage()
			return
		}
		factory197 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt198 := factory197.GetProtocol(mbTrans195)
		argvalue0 := api.NewJobKey()
		err199 := argvalue0.Read(jsProt198)
		if err199 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.StartCronJob(value0))
		fmt.Print("\n")
		break
	case "restartShards":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "RestartShards requires 3 args")
			flag.Usage()
		}
		arg200 := flag.Arg(1)
		mbTrans201 := thrift.NewTMemoryBufferLen(len(arg200))
		defer mbTrans201.Close()
		_, err202 := mbTrans201.WriteString(arg200)
		if err202 != nil {
			Usage()
			return
		}
		factory203 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt204 := factory203.GetProtocol(mbTrans201)
		argvalue0 := api.NewJobKey()
		err205 := argvalue0.Read(jsProt204)
		if err205 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg206 := flag.Arg(2)
		mbTrans207 := thrift.NewTMemoryBufferLen(len(arg206))
		defer mbTrans207.Close()
		_, err208 := mbTrans207.WriteString(arg206)
		if err208 != nil {
			Usage()
			return
		}
		factory209 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt210 := factory209.GetProtocol(mbTrans207)
		containerStruct1 := api.NewAuroraSchedulerManagerRestartShardsArgs()
		err211 := containerStruct1.ReadField2(jsProt210)
		if err211 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.ShardIds
		value1 := argvalue1
		arg212 := flag.Arg(3)
		mbTrans213 := thrift.NewTMemoryBufferLen(len(arg212))
		defer mbTrans213.Close()
		_, err214 := mbTrans213.WriteString(arg212)
		if err214 != nil {
			Usage()
			return
		}
		factory215 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt216 := factory215.GetProtocol(mbTrans213)
		argvalue2 := api.NewLock()
		err217 := argvalue2.Read(jsProt216)
		if err217 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.RestartShards(value0, value1, value2))
		fmt.Print("\n")
		break
	case "killTasks":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "KillTasks requires 4 args")
			flag.Usage()
		}
		arg218 := flag.Arg(1)
		mbTrans219 := thrift.NewTMemoryBufferLen(len(arg218))
		defer mbTrans219.Close()
		_, err220 := mbTrans219.WriteString(arg218)
		if err220 != nil {
			Usage()
			return
		}
		factory221 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt222 := factory221.GetProtocol(mbTrans219)
		argvalue0 := api.NewTaskQuery()
		err223 := argvalue0.Read(jsProt222)
		if err223 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg224 := flag.Arg(2)
		mbTrans225 := thrift.NewTMemoryBufferLen(len(arg224))
		defer mbTrans225.Close()
		_, err226 := mbTrans225.WriteString(arg224)
		if err226 != nil {
			Usage()
			return
		}
		factory227 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt228 := factory227.GetProtocol(mbTrans225)
		argvalue1 := api.NewLock()
		err229 := argvalue1.Read(jsProt228)
		if err229 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		arg230 := flag.Arg(3)
		mbTrans231 := thrift.NewTMemoryBufferLen(len(arg230))
		defer mbTrans231.Close()
		_, err232 := mbTrans231.WriteString(arg230)
		if err232 != nil {
			Usage()
			return
		}
		factory233 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt234 := factory233.GetProtocol(mbTrans231)
		argvalue2 := api.NewJobKey()
		err235 := argvalue2.Read(jsProt234)
		if err235 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		arg236 := flag.Arg(4)
		mbTrans237 := thrift.NewTMemoryBufferLen(len(arg236))
		defer mbTrans237.Close()
		_, err238 := mbTrans237.WriteString(arg236)
		if err238 != nil {
			Usage()
			return
		}
		factory239 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt240 := factory239.GetProtocol(mbTrans237)
		containerStruct3 := api.NewAuroraSchedulerManagerKillTasksArgs()
		err241 := containerStruct3.ReadField4(jsProt240)
		if err241 != nil {
			Usage()
			return
		}
		argvalue3 := containerStruct3.Instances
		value3 := argvalue3
		fmt.Print(client.KillTasks(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "addInstances":
		if flag.NArg()-1 != 4 {
			fmt.Fprintln(os.Stderr, "AddInstances requires 4 args")
			flag.Usage()
		}
		arg242 := flag.Arg(1)
		mbTrans243 := thrift.NewTMemoryBufferLen(len(arg242))
		defer mbTrans243.Close()
		_, err244 := mbTrans243.WriteString(arg242)
		if err244 != nil {
			Usage()
			return
		}
		factory245 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt246 := factory245.GetProtocol(mbTrans243)
		argvalue0 := api.NewAddInstancesConfig()
		err247 := argvalue0.Read(jsProt246)
		if err247 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg248 := flag.Arg(2)
		mbTrans249 := thrift.NewTMemoryBufferLen(len(arg248))
		defer mbTrans249.Close()
		_, err250 := mbTrans249.WriteString(arg248)
		if err250 != nil {
			Usage()
			return
		}
		factory251 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt252 := factory251.GetProtocol(mbTrans249)
		argvalue1 := api.NewLock()
		err253 := argvalue1.Read(jsProt252)
		if err253 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		arg254 := flag.Arg(3)
		mbTrans255 := thrift.NewTMemoryBufferLen(len(arg254))
		defer mbTrans255.Close()
		_, err256 := mbTrans255.WriteString(arg254)
		if err256 != nil {
			Usage()
			return
		}
		factory257 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt258 := factory257.GetProtocol(mbTrans255)
		argvalue2 := api.NewInstanceKey()
		err259 := argvalue2.Read(jsProt258)
		if err259 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		tmp3, err260 := (strconv.Atoi(flag.Arg(4)))
		if err260 != nil {
			Usage()
			return
		}
		argvalue3 := int32(tmp3)
		value3 := argvalue3
		fmt.Print(client.AddInstances(value0, value1, value2, value3))
		fmt.Print("\n")
		break
	case "acquireLock":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "AcquireLock requires 1 args")
			flag.Usage()
		}
		arg261 := flag.Arg(1)
		mbTrans262 := thrift.NewTMemoryBufferLen(len(arg261))
		defer mbTrans262.Close()
		_, err263 := mbTrans262.WriteString(arg261)
		if err263 != nil {
			Usage()
			return
		}
		factory264 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt265 := factory264.GetProtocol(mbTrans262)
		argvalue0 := api.NewLockKey()
		err266 := argvalue0.Read(jsProt265)
		if err266 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.AcquireLock(value0))
		fmt.Print("\n")
		break
	case "releaseLock":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ReleaseLock requires 2 args")
			flag.Usage()
		}
		arg267 := flag.Arg(1)
		mbTrans268 := thrift.NewTMemoryBufferLen(len(arg267))
		defer mbTrans268.Close()
		_, err269 := mbTrans268.WriteString(arg267)
		if err269 != nil {
			Usage()
			return
		}
		factory270 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt271 := factory270.GetProtocol(mbTrans268)
		argvalue0 := api.NewLock()
		err272 := argvalue0.Read(jsProt271)
		if err272 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err := (strconv.Atoi(flag.Arg(2)))
		if err != nil {
			Usage()
			return
		}
		argvalue1 := api.LockValidation(tmp1)
		value1 := argvalue1
		fmt.Print(client.ReleaseLock(value0, value1))
		fmt.Print("\n")
		break
	case "replaceCronTemplate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ReplaceCronTemplate requires 2 args")
			flag.Usage()
		}
		arg273 := flag.Arg(1)
		mbTrans274 := thrift.NewTMemoryBufferLen(len(arg273))
		defer mbTrans274.Close()
		_, err275 := mbTrans274.WriteString(arg273)
		if err275 != nil {
			Usage()
			return
		}
		factory276 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt277 := factory276.GetProtocol(mbTrans274)
		argvalue0 := api.NewJobConfiguration()
		err278 := argvalue0.Read(jsProt277)
		if err278 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg279 := flag.Arg(2)
		mbTrans280 := thrift.NewTMemoryBufferLen(len(arg279))
		defer mbTrans280.Close()
		_, err281 := mbTrans280.WriteString(arg279)
		if err281 != nil {
			Usage()
			return
		}
		factory282 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt283 := factory282.GetProtocol(mbTrans280)
		argvalue1 := api.NewLock()
		err284 := argvalue1.Read(jsProt283)
		if err284 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.ReplaceCronTemplate(value0, value1))
		fmt.Print("\n")
		break
	case "startJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "StartJobUpdate requires 2 args")
			flag.Usage()
		}
		arg285 := flag.Arg(1)
		mbTrans286 := thrift.NewTMemoryBufferLen(len(arg285))
		defer mbTrans286.Close()
		_, err287 := mbTrans286.WriteString(arg285)
		if err287 != nil {
			Usage()
			return
		}
		factory288 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt289 := factory288.GetProtocol(mbTrans286)
		argvalue0 := api.NewJobUpdateRequest()
		err290 := argvalue0.Read(jsProt289)
		if err290 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.StartJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "pauseJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PauseJobUpdate requires 2 args")
			flag.Usage()
		}
		arg292 := flag.Arg(1)
		mbTrans293 := thrift.NewTMemoryBufferLen(len(arg292))
		defer mbTrans293.Close()
		_, err294 := mbTrans293.WriteString(arg292)
		if err294 != nil {
			Usage()
			return
		}
		factory295 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt296 := factory295.GetProtocol(mbTrans293)
		argvalue0 := api.NewJobUpdateKey()
		err297 := argvalue0.Read(jsProt296)
		if err297 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.PauseJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "resumeJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ResumeJobUpdate requires 2 args")
			flag.Usage()
		}
		arg299 := flag.Arg(1)
		mbTrans300 := thrift.NewTMemoryBufferLen(len(arg299))
		defer mbTrans300.Close()
		_, err301 := mbTrans300.WriteString(arg299)
		if err301 != nil {
			Usage()
			return
		}
		factory302 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt303 := factory302.GetProtocol(mbTrans300)
		argvalue0 := api.NewJobUpdateKey()
		err304 := argvalue0.Read(jsProt303)
		if err304 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.ResumeJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "abortJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AbortJobUpdate requires 2 args")
			flag.Usage()
		}
		arg306 := flag.Arg(1)
		mbTrans307 := thrift.NewTMemoryBufferLen(len(arg306))
		defer mbTrans307.Close()
		_, err308 := mbTrans307.WriteString(arg306)
		if err308 != nil {
			Usage()
			return
		}
		factory309 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt310 := factory309.GetProtocol(mbTrans307)
		argvalue0 := api.NewJobUpdateKey()
		err311 := argvalue0.Read(jsProt310)
		if err311 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.AbortJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "pulseJobUpdate":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PulseJobUpdate requires 1 args")
			flag.Usage()
		}
		arg313 := flag.Arg(1)
		mbTrans314 := thrift.NewTMemoryBufferLen(len(arg313))
		defer mbTrans314.Close()
		_, err315 := mbTrans314.WriteString(arg313)
		if err315 != nil {
			Usage()
			return
		}
		factory316 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt317 := factory316.GetProtocol(mbTrans314)
		argvalue0 := api.NewJobUpdateKey()
		err318 := argvalue0.Read(jsProt317)
		if err318 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PulseJobUpdate(value0))
		fmt.Print("\n")
		break
	case "getRoleSummary":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetRoleSummary requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetRoleSummary())
		fmt.Print("\n")
		break
	case "getJobSummary":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobSummary requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetJobSummary(value0))
		fmt.Print("\n")
		break
	case "getTasksStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTasksStatus requires 1 args")
			flag.Usage()
		}
		arg320 := flag.Arg(1)
		mbTrans321 := thrift.NewTMemoryBufferLen(len(arg320))
		defer mbTrans321.Close()
		_, err322 := mbTrans321.WriteString(arg320)
		if err322 != nil {
			Usage()
			return
		}
		factory323 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt324 := factory323.GetProtocol(mbTrans321)
		argvalue0 := api.NewTaskQuery()
		err325 := argvalue0.Read(jsProt324)
		if err325 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTasksStatus(value0))
		fmt.Print("\n")
		break
	case "getTasksWithoutConfigs":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTasksWithoutConfigs requires 1 args")
			flag.Usage()
		}
		arg326 := flag.Arg(1)
		mbTrans327 := thrift.NewTMemoryBufferLen(len(arg326))
		defer mbTrans327.Close()
		_, err328 := mbTrans327.WriteString(arg326)
		if err328 != nil {
			Usage()
			return
		}
		factory329 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt330 := factory329.GetProtocol(mbTrans327)
		argvalue0 := api.NewTaskQuery()
		err331 := argvalue0.Read(jsProt330)
		if err331 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTasksWithoutConfigs(value0))
		fmt.Print("\n")
		break
	case "getPendingReason":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPendingReason requires 1 args")
			flag.Usage()
		}
		arg332 := flag.Arg(1)
		mbTrans333 := thrift.NewTMemoryBufferLen(len(arg332))
		defer mbTrans333.Close()
		_, err334 := mbTrans333.WriteString(arg332)
		if err334 != nil {
			Usage()
			return
		}
		factory335 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt336 := factory335.GetProtocol(mbTrans333)
		argvalue0 := api.NewTaskQuery()
		err337 := argvalue0.Read(jsProt336)
		if err337 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetPendingReason(value0))
		fmt.Print("\n")
		break
	case "getConfigSummary":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetConfigSummary requires 1 args")
			flag.Usage()
		}
		arg338 := flag.Arg(1)
		mbTrans339 := thrift.NewTMemoryBufferLen(len(arg338))
		defer mbTrans339.Close()
		_, err340 := mbTrans339.WriteString(arg338)
		if err340 != nil {
			Usage()
			return
		}
		factory341 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt342 := factory341.GetProtocol(mbTrans339)
		argvalue0 := api.NewJobKey()
		err343 := argvalue0.Read(jsProt342)
		if err343 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetConfigSummary(value0))
		fmt.Print("\n")
		break
	case "getJobs":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobs requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetJobs(value0))
		fmt.Print("\n")
		break
	case "getQuota":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetQuota requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetQuota(value0))
		fmt.Print("\n")
		break
	case "populateJobConfig":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PopulateJobConfig requires 1 args")
			flag.Usage()
		}
		arg346 := flag.Arg(1)
		mbTrans347 := thrift.NewTMemoryBufferLen(len(arg346))
		defer mbTrans347.Close()
		_, err348 := mbTrans347.WriteString(arg346)
		if err348 != nil {
			Usage()
			return
		}
		factory349 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt350 := factory349.GetProtocol(mbTrans347)
		argvalue0 := api.NewJobConfiguration()
		err351 := argvalue0.Read(jsProt350)
		if err351 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PopulateJobConfig(value0))
		fmt.Print("\n")
		break
	case "getLocks":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetLocks requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetLocks())
		fmt.Print("\n")
		break
	case "getJobUpdateSummaries":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateSummaries requires 1 args")
			flag.Usage()
		}
		arg352 := flag.Arg(1)
		mbTrans353 := thrift.NewTMemoryBufferLen(len(arg352))
		defer mbTrans353.Close()
		_, err354 := mbTrans353.WriteString(arg352)
		if err354 != nil {
			Usage()
			return
		}
		factory355 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt356 := factory355.GetProtocol(mbTrans353)
		argvalue0 := api.NewJobUpdateQuery()
		err357 := argvalue0.Read(jsProt356)
		if err357 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateSummaries(value0))
		fmt.Print("\n")
		break
	case "getJobUpdateDetails":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateDetails requires 1 args")
			flag.Usage()
		}
		arg358 := flag.Arg(1)
		mbTrans359 := thrift.NewTMemoryBufferLen(len(arg358))
		defer mbTrans359.Close()
		_, err360 := mbTrans359.WriteString(arg358)
		if err360 != nil {
			Usage()
			return
		}
		factory361 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt362 := factory361.GetProtocol(mbTrans359)
		argvalue0 := api.NewJobUpdateKey()
		err363 := argvalue0.Read(jsProt362)
		if err363 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateDetails(value0))
		fmt.Print("\n")
		break
	case "getJobUpdateDiff":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateDiff requires 1 args")
			flag.Usage()
		}
		arg364 := flag.Arg(1)
		mbTrans365 := thrift.NewTMemoryBufferLen(len(arg364))
		defer mbTrans365.Close()
		_, err366 := mbTrans365.WriteString(arg364)
		if err366 != nil {
			Usage()
			return
		}
		factory367 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt368 := factory367.GetProtocol(mbTrans365)
		argvalue0 := api.NewJobUpdateRequest()
		err369 := argvalue0.Read(jsProt368)
		if err369 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateDiff(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
