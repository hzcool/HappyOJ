(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0c8a86fd"],{"394ca":function(t,e,a){"use strict";a.r(e);var o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return t.loaded?a("div",[a("Card",[a("h1",[t._v("说明")]),a("h2",[t._v("1.仅支持"),a("span",{staticStyle:{color:"red"}},[t._v("zip压缩包")]),t._v("文件")]),a("h2",[t._v("2.多次导入,前面导入的文件会被删除")]),a("h2",[t._v("3.输入输出可以文件名应该匹配,否则无法失败; 合法的输入文件后缀是.in，输出文件的后缀是.ans或者.out;输入输出文件名应该匹配,比如(1.in,1.out)、(data01.in.data01.ans)等, 提交到后台进行处理并重命名为测试号加后缀的文件格式")]),a("br"),a("Divider"),a("Row",[a("i-col",{attrs:{span:"4"}},[a("Upload",{attrs:{"before-upload":t.before_upload,type:"drag",action:"//jsonplaceholder.typicode.com/posts/"}},[a("div",{staticStyle:{padding:"20px 0"}},[a("Icon",{staticStyle:{color:"#3399ff"},attrs:{type:"ios-cloud-upload",size:"52"}}),a("p",[t._v("选择文件")])],1)])],1)],1),t.choose_file?a("div",[t._v(" "+t._s(t.choose_file.name)+" "),a("Button",{staticStyle:{"margin-left":"20px"},attrs:{type:"success"},on:{click:t.submit}},[t._v(" 上传 ")]),a("Button",{staticStyle:{"margin-left":"20px"},attrs:{type:"error"},on:{click:t.remove_file}},[t._v(" 移除 ")])],1):t._e(),a("Divider"),t.data.length?a("div",[a("span",{staticStyle:{"font-size":"20px",color:"black"}},[a("strong",[t._v("全部测试数据如下")])]),a("Button",{staticStyle:{float:"right"},attrs:{type:"error"},on:{click:t.delFile}},[t._v(" 删除全部 ")]),a("div",{staticStyle:{clear:"both","margin-top":"10px"}}),a("Table",{attrs:{border:"",columns:t.columns,data:t.data},scopedSlots:t._u([{key:"action",fn:function(e){var o=e.row,i=e.index;return[a("Button",{staticStyle:{"margin-right":"5px"},attrs:{type:"primary",size:"small",loading:t.loading},on:{click:function(e){return t.showTestdata(o)}}},[t._v("查看")]),a("Button",{attrs:{type:"error",size:"small"},on:{click:function(e){return t.removeOneTestdata(o,i)}}},[t._v("删除")])]}}],null,!1,692043013)})],1):t._e(),a("Modal",{staticStyle:{width:"800px"},attrs:{"mask-closable":!1,draggable:"",scrollable:""},on:{"on-ok":t.ok,"on-cancel":t.ok},model:{value:t.modal,callback:function(e){t.modal=e},expression:"modal"}},[a("p",{staticStyle:{color:"red"}},[t._v(" "+t._s(t.modal_input_name)+" ")]),a("Input",{attrs:{type:"textarea",autosize:{minRows:2,maxRows:10}},model:{value:t.modal_input,callback:function(e){t.modal_input=e},expression:"modal_input"}}),a("Divider"),a("p",{staticStyle:{color:"red"}},[t._v(" "+t._s(t.modal_output_name)+" ")]),a("Input",{attrs:{type:"textarea",autosize:{minRows:2,maxRows:10}},model:{value:t.modal_output,callback:function(e){t.modal_output=e},expression:"modal_output"}})],1)],1)],1):t._e()},i=[],n=(a("a434"),{data:function(){return{loaded:!1,index:"",choose_file:null,columns:[{title:"测试点",key:"case_id"},{title:"输入文件",key:"input_name"},{title:"输入文件大小",key:"input_size"},{title:"输出文件",key:"output_name"},{title:"输出文件大小",key:"output_size"},{title:"操作",slot:"action"}],data:[],modal:!1,modal_input_name:"",modal_input:"",modal_output_name:"",modal_output:"",loading:!1}},methods:{before_upload:function(t){return this.choose_file=t,!1},delFile:function(){var t=this;this.$req.get({url:"removeTestdatas",params:{index:this.index}}).then((function(e){e&&(t.$message("success",e),t.data=[])}))},submit:function(){var t=this,e=new FormData,a={headers:{"Content-Type":"multipart/form-data"}};e.append("index",this.index),e.append("zip",this.choose_file),this.$req.post({url:"addTestdatas",form:e,config:a}).then((function(e){e&&(t.choose_file=null,t.$message("success",e),t.getTestdatas())}))},remove_file:function(){this.choose_file=null},showTestdata:function(t){var e=this;this.loading=!0,this.$req.get({url:"getOneTestdata",params:{index:this.index,case_id:t.case_id}}).then((function(a){e.loading=!1,a&&(e.modal_input_name=t.input_name,e.modal_output_name=t.output_name,e.modal_input=a.input,e.modal_output=a.output,e.modal=!0)}))},removeOneTestdata:function(t,e){var a=this;this.$req.get({url:"removeOneTestdata",params:{index:this.index,case_id:t.case_id}}).then((function(t){t&&(a.$message("success",t),a.data.splice(e,1))}))},getTestdatas:function(){var t=this;this.$req.get({url:"getTestdatas",params:{index:this.index}}).then((function(e){e.data&&(t.data=e.data)}))},ok:function(){this.modal=!1}},mounted:function(){this.$route.query.index?(this.index=this.$route.query.index,this.getTestdatas(),this.loaded=!0):(this.$message("error","出错"),this.$router.push({name:"problemManage"}))}}),s=n,l=a("2877"),r=Object(l["a"])(s,o,i,!1,null,null,null);e["default"]=r.exports},a434:function(t,e,a){"use strict";var o=a("23e7"),i=a("23cb"),n=a("a691"),s=a("50c4"),l=a("7b0b"),r=a("65f0"),d=a("8418"),u=a("1dde"),c=a("ae40"),p=u("splice"),m=c("splice",{ACCESSORS:!0,0:0,1:2}),_=Math.max,h=Math.min,f=9007199254740991,v="Maximum allowed length exceeded";o({target:"Array",proto:!0,forced:!p||!m},{splice:function(t,e){var a,o,u,c,p,m,g=l(this),y=s(g.length),x=i(t,y),b=arguments.length;if(0===b?a=o=0:1===b?(a=0,o=y-x):(a=b-2,o=h(_(n(e),0),y-x)),y+a-o>f)throw TypeError(v);for(u=r(g,o),c=0;c<o;c++)p=x+c,p in g&&d(u,c,g[p]);if(u.length=o,a<o){for(c=x;c<y-o;c++)p=c+o,m=c+a,p in g?g[m]=g[p]:delete g[m];for(c=y;c>y-o+a;c--)delete g[c-1]}else if(a>o)for(c=y-o;c>x;c--)p=c+o-1,m=c+a-1,p in g?g[m]=g[p]:delete g[m];for(c=0;c<a;c++)g[c+x]=arguments[c+2];return g.length=y-o+a,u}})}}]);