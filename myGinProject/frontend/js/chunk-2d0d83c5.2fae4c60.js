(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0d83c5"],{"79e2":function(t,e,s){"use strict";s.r(e);var a=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("Card",[s("List",{staticStyle:{float:"left",color:"black"},attrs:{border:"",size:"large"}},[s("ListItem",[s("Row",{staticStyle:{width:"800px"}},[s("i-col",{attrs:{span:"4"}},[t._v("用户名")]),s("i-col",{staticStyle:{color:"red"},attrs:{span:"20"}},[t._v(t._s(t.info.username))])],1)],1),s("ListItem",[s("Row",{staticStyle:{width:"800px"}},[s("i-col",{attrs:{span:"4"}},[t._v("注册时间")]),s("i-col",{staticStyle:{color:"red"},attrs:{span:"20"}},[t._v(t._s(t.info.create_at))])],1)],1),s("ListItem",[s("Row",{staticStyle:{width:"800px"}},[s("i-col",{attrs:{span:"4"}},[t._v("上次访问")]),s("i-col",{staticStyle:{color:"red"},attrs:{span:"20"}},[t._v(t._s(t.info.last_seen))])],1)],1),s("ListItem",[s("Row",{staticStyle:{width:"800px"}},[s("i-col",{attrs:{span:"4"}},[t._v("学校")]),s("i-col",{staticStyle:{color:"red"},attrs:{span:"20"}},[t._v(t._s(t.info.school))])],1)],1),s("ListItem",[s("Row",{staticStyle:{width:"800px"}},[s("i-col",{attrs:{span:"4"}},[t._v("邮箱")]),s("i-col",{staticStyle:{color:"red"},attrs:{span:"20"}},[t._v(t._s(t.info.email))])],1)],1),s("ListItem",[s("Row",{staticStyle:{width:"800px"}},[s("i-col",{attrs:{span:"4"}},[t._v("描述")]),s("i-col",{staticStyle:{color:"red"},attrs:{span:"20"}},[t._v(t._s(t.info.desc))])],1)],1),s("ListItem",[s("Row",{staticStyle:{width:"800px"}},[s("i-col",{attrs:{span:"4"}},[t._v(" 已解决的问题("+t._s(t.solved.length)+") ")]),s("i-col",{staticStyle:{color:"blue"},attrs:{span:"20"}},t._l(t.solved,(function(e){return s("router-link",{key:e,attrs:{to:{name:"problem",params:{Index:e}}}},[t._v(" "+t._s(e)+" ")])})),1)],1)],1),s("ListItem",[s("Row",{staticStyle:{width:"800px"}},[s("i-col",{attrs:{span:"4"}},[t._v(" 未解决的问题("+t._s(t.unsolved.length)+") ")]),s("i-col",{staticStyle:{color:"blue"},attrs:{span:"20"}},t._l(t.unsolved,(function(e){return s("router-link",{key:e,attrs:{to:{name:"problem",params:{Index:e}}}},[t._v(" "+t._s(e)+" ")])})),1)],1)],1)],1),s("div",{staticClass:"myChart",style:{width:"700px",height:"600px",margin:"0 auto",float:"right"},attrs:{id:"myChart"}}),s("div",{staticStyle:{clear:"both"}})],1)},i=[],o={data:function(){return{info:{username:"hzcool",school:"南京航空航天大学",email:"562954019@qq.com",desc:"我是个憨批",create_at:"gg",last_seen:"gg"},solved:[],mp:{Accepted:"正确",WrongAnswer:"答案错误",TimeLimitExceeded:"时间超限",MemoryLimitExceeded:"空间超限",OutputLimitExceeded:"输出超限",RuntimeError:"运行错误",SystemError:"系统错误",CompileError:"编译错误"},unsolved:[],pie:{title:{text:"解题分析",left:"center"},tooltip:{trigger:"item",formatter:"{a} <br/>{b} : {c} ({d}%)"},legend:{orient:"vertical",left:"left",data:["正确","答案错误","时间超限","空间超限","输出超限","运行错误","系统错误","编译错误"]},series:[{name:"解题分析",type:"pie",radius:"55%",center:["50%","60%"],data:[],emphasis:{itemStyle:{shadowBlur:10,shadowOffsetX:0,shadowColor:"rgba(0, 0, 0, 0.5)"}},itemStyle:{normal:{color:function(t){var e=["#00FF00","#A52A2A","#95DE64","#5CDBD3","#69C0FF","#FFBB96","DeepPink","red"];return e[t.dataIndex]}}}}]}}},methods:{drawLine:function(){var t=this.$echarts.init(document.getElementById("myChart"));t.setOption(this.pie)}},mounted:function(){var t=this,e=this.$route.query.username;e?this.$req.get({url:"getUserInfo",params:{username:e}}).then((function(e){if(e){for(var s in t.info=e.info,e.solved&&(t.solved=e.solved),e.unsolved&&(t.unsolved=e.unsolved),e.status)t.pie.series[0].data.push({name:t.mp[s],value:e.status[s]});t.drawLine()}})):(alert("出错"),this.$router.push({name:"home"}))}},r=o,l=s("2877"),n=Object(l["a"])(r,a,i,!1,null,null,null);e["default"]=n.exports}}]);