(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-ce262f38"],{1148:function(t,e,o){"use strict";var r=o("a691"),n=o("1d80");t.exports="".repeat||function(t){var e=String(n(this)),o="",s=r(t);if(s<0||s==1/0)throw RangeError("Wrong number of repetitions");for(;s>0;(s>>>=1)&&(e+=e))1&s&&(o+=e);return o}},"129f":function(t,e){t.exports=Object.is||function(t,e){return t===e?0!==t||1/t===1/e:t!=t&&e!=e}},"14c3":function(t,e,o){var r=o("c6b6"),n=o("9263");t.exports=function(t,e){var o=t.exec;if("function"===typeof o){var s=o.call(t,e);if("object"!==typeof s)throw TypeError("RegExp exec method returned something other than an Object or null");return s}if("RegExp"!==r(t))throw TypeError("RegExp#exec called on incompatible receiver");return n.call(t,e)}},"2e8c":function(t,e,o){"use strict";o.r(e);var r=function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",{style:t.backgroundDiv},[o("Affix",{attrs:{"offset-top":50}},[o("div",{staticClass:"block"},[o("h2",{staticStyle:{color:"#fff"}},[t._v("Happy Online Judge")]),o("Divider",{attrs:{size:"small"}}),0==t.type?o("div",[o("Form",{ref:"info",attrs:{model:t.info}},[o("h2",{staticStyle:{color:"#fff"}},[t._v("修改信息")]),t._v(" "),o("br"),o("FormItem",[o("Row",[o("Col",{staticStyle:{color:"#fff","font-size":"20px"},attrs:{span:"4"}},[o("span",{staticStyle:{color:"red"}},[t._v("*")]),t._v("用户名 ")]),o("Col",{attrs:{span:"18",offset:"1"}},[o("Input",{attrs:{type:"text",placeholder:"Username",size:"large",autofocus:"",clearable:""},model:{value:t.info.username,callback:function(e){t.$set(t.info,"username",e)},expression:"info.username"}},[o("Icon",{attrs:{slot:"prepend",type:"ios-person-outline"},slot:"prepend"})],1)],1)],1)],1),o("FormItem",[o("Row",[o("Col",{staticStyle:{color:"#fff","font-size":"20px"},attrs:{span:"4"}},[o("span",{staticStyle:{color:"red"}},[t._v("*")]),t._v("密 码 ")]),o("Col",{attrs:{span:"18",offset:"1"}},[o("Input",{attrs:{type:"password",placeholder:"Password",size:"large",clearable:""},model:{value:t.info.password,callback:function(e){t.$set(t.info,"password",e)},expression:"info.password"}},[o("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1)],1)],1),o("FormItem",[o("Row",[o("Col",{staticStyle:{color:"#fff","font-size":"20px"},attrs:{span:"4"}},[o("span",{staticStyle:{color:"red"}},[t._v("*")]),t._v("邮 箱 ")]),o("Col",{attrs:{span:"18",offset:"1"}},[o("Input",{attrs:{type:"text",disabled:"",placeholder:"Email",size:"large",clearable:""},model:{value:t.info.email,callback:function(e){t.$set(t.info,"email",e)},expression:"info.email"}},[o("Icon",{attrs:{slot:"prepend",type:"ios-mail-outline"},slot:"prepend"})],1)],1)],1)],1),o("FormItem",[o("Row",[o("Col",{staticStyle:{color:"#fff","font-size":"20px"},attrs:{span:"4"}},[o("span",{staticStyle:{color:"red"}},[t._v("*")]),t._v("学 校 ")]),o("Col",{attrs:{span:"18",offset:"1"}},[o("Input",{attrs:{type:"text",placeholder:"School",size:"large",clearable:""},model:{value:t.info.school,callback:function(e){t.$set(t.info,"school",e)},expression:"info.school"}},[o("Icon",{attrs:{slot:"prepend",type:"ios-school-outline"},slot:"prepend"})],1)],1)],1)],1),o("FormItem",[o("Col",{staticStyle:{color:"#fff","font-size":"20px"},attrs:{span:"4"}},[t._v(" 描 述 ")]),o("Col",{attrs:{span:"18",offset:"1"}},[o("Input",{attrs:{type:"textarea",rows:4,placeholder:"Enter something...",clearable:""},model:{value:t.info.description,callback:function(e){t.$set(t.info,"description",e)},expression:"info.description"}})],1)],1),o("br"),o("FormItem",[o("Row",{attrs:{type:"flex",justify:"center"}},[o("Col",{attrs:{span:"5"}},[o("Button",{attrs:{type:"primary",size:"large"},on:{click:t.submit}},[t._v("提交")])],1),o("Col",{attrs:{span:"5"}},[o("Button",{attrs:{type:"error",size:"large"},on:{click:t.clear}},[t._v("清空")])],1)],1)],1),o("br")],1),o("Row",{staticStyle:{"font-size":"20px",color:"#ddd"},attrs:{type:"flex",justify:"end"}},[o("Button",{staticStyle:{color:"red","margin-right":"10px"},attrs:{type:"dashed"},on:{click:t.changeType}},[t._v(" 修改密码 ")]),o("router-link",{attrs:{to:{name:"home"}}},[t._v(" 返回主页 ")])],1)],1):o("div",[o("Form",[o("h2",{staticStyle:{color:"#fff"}},[t._v("修改密码")]),t._v(" "),o("br"),o("FormItem",[o("Row",[o("Col",{staticStyle:{color:"#fff","font-size":"20px"},attrs:{span:"5"}},[o("span",{staticStyle:{color:"red"}},[t._v("*")]),t._v("旧密码 ")]),o("Col",{attrs:{span:"18",offset:"1"}},[o("Input",{attrs:{type:"password",placeholder:"old password",size:"large",clearable:""},model:{value:t.pwd.oldPassword,callback:function(e){t.$set(t.pwd,"oldPassword",e)},expression:"pwd.oldPassword"}},[o("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1)],1)],1),o("FormItem",[o("Row",[o("Col",{staticStyle:{color:"#fff","font-size":"20px"},attrs:{span:"5"}},[o("span",{staticStyle:{color:"red"}},[t._v("*")]),t._v("新密码 ")]),o("Col",{attrs:{span:"18",offset:"1"}},[o("Input",{attrs:{type:"password",placeholder:"new password",size:"large",clearable:""},model:{value:t.pwd.newPassword,callback:function(e){t.$set(t.pwd,"newPassword",e)},expression:"pwd.newPassword"}},[o("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1)],1)],1),o("FormItem",[o("Row",[o("Col",{staticStyle:{color:"#fff","font-size":"20px"},attrs:{span:"5"}},[o("span",{staticStyle:{color:"red"}},[t._v("*")]),t._v("重复新密码 ")]),o("Col",{attrs:{span:"18",offset:"1"}},[o("Input",{attrs:{type:"password",placeholder:"repeat",size:"large",clearable:""},model:{value:t.repeat,callback:function(e){t.repeat=e},expression:"repeat"}},[o("Icon",{attrs:{slot:"prepend",type:"ios-lock-outline"},slot:"prepend"})],1)],1)],1)],1),o("FormItem",[o("Row",{attrs:{type:"flex",justify:"center"}},[o("Col",{attrs:{span:"5"}},[o("Button",{attrs:{type:"primary",size:"large"},on:{click:t.updatePassword}},[t._v("提交")])],1),o("Col",{attrs:{span:"5"}},[o("Button",{attrs:{type:"error",size:"large"},on:{click:t.clear2}},[t._v("清空")])],1)],1)],1)],1),o("Row",{staticStyle:{"font-size":"20px",color:"#ddd"},attrs:{type:"flex",justify:"end"}},[o("Button",{staticStyle:{color:"red","margin-right":"10px"},attrs:{type:"dashed"},on:{click:t.changeType}},[t._v(" 修改信息 ")]),o("router-link",{attrs:{to:{name:"home"}}},[t._v(" 返回主页 ")])],1)],1)],1)])],1)},n=[],s=(o("a4d3"),o("e01a"),o("ac1f"),o("38cf"),o("841c"),{name:"login",data:function(){return{backgroundDiv:{backgroundImage:"url("+o("8554")+")",backgroundRepeat:"no-repeat",backgroundSize:"100% 100%",height:"1100px"},pwd:{oldPassword:"",newPassword:""},repeat:"",type:0,info:{username:"",password:"",email:"",school:"",description:""}}},methods:{validate:function(t){return 0==t.username.length?"用户名不能为空":t.username.length>20?"用户名太长":-1!=t.username.search(" ")?"用户名不能包含空格":0==t.password.length?"密码不能为空":t.password.length>16?"密码太长":-1!=t.password.search(" ")?"密码不能包含空格":0==t.school.length?"学校不能为空":t.school.length>20?"学校名太长":-1!=t.school.search(" ")?"学校名不能包含空格":""},submit:function(){var t=this,e=this.validate(this.info);""!=e?this.$message("error",e):this.$req.post({url:"update",form:JSON.stringify(this.info),config:{headers:{"Content-Type":" application/json"}}}).then((function(e){200!=e.status?t.$message("error",e.data):(t.$message("success",e.data),t.$store.commit("Login",t.info.username))}))},changeType:function(){this.type^=1},clear:function(){this.info.description=this.info.school=this.info.email=this.info.username=this.info.password=""},validate2:function(t){return 0==t.oldPassword.length?"旧密码不能为空":t.oldPassword.length<6?"旧密码太短":t.oldPassword.length>16?"旧密码太长":t.oldPassword==t.newPassword?"新密码不能与旧密码相同":0==t.newPassword.length?"新密码不能为空":t.newPassword.length<6?"新密码太短":t.newPassword.length>16?"新密码太长":0==this.repeat.length?"请重复新密码":t.newPassword!=this.repeat?"两次密码输入不同":""},updatePassword:function(){var t=this,e=this.validate2(this.pwd);""!=e?this.$message("error",e):this.$req.post({url:"updatePassword",form:JSON.stringify(this.pwd),config:{headers:{"Content-Type":" application/json"}}}).then((function(e){e&&t.$message("success",e.data)}))},clear2:function(){this.repeat=this.pwd.newPassword=this.pwd.oldPassword=""}},mounted:function(){var t=this;this.$req.get({url:"profile"}).then((function(e){e&&(t.info.username=e.username,t.info.email=e.email,t.info.school=e.school,t.info.description=e.description)}))}}),a=s,i=(o("505c"),o("2877")),l=Object(i["a"])(a,r,n,!1,null,"2a739ebc",null);e["default"]=l.exports},"38cf":function(t,e,o){var r=o("23e7"),n=o("1148");r({target:"String",proto:!0},{repeat:n})},"505c":function(t,e,o){"use strict";var r=o("f0c8"),n=o.n(r);n.a},"841c":function(t,e,o){"use strict";var r=o("d784"),n=o("825a"),s=o("1d80"),a=o("129f"),i=o("14c3");r("search",1,(function(t,e,o){return[function(e){var o=s(this),r=void 0==e?void 0:e[t];return void 0!==r?r.call(e,o):new RegExp(e)[t](String(o))},function(t){var r=o(e,t,this);if(r.done)return r.value;var s=n(t),l=String(this),c=s.lastIndex;a(c,0)||(s.lastIndex=0);var p=i(s,l);return a(s.lastIndex,c)||(s.lastIndex=c),null===p?-1:p.index}]}))},8554:function(t,e,o){t.exports=o.p+"img/1.b8056ce3.jpg"},9263:function(t,e,o){"use strict";var r=o("ad6d"),n=o("9f7fd"),s=RegExp.prototype.exec,a=String.prototype.replace,i=s,l=function(){var t=/a/,e=/b*/g;return s.call(t,"a"),s.call(e,"a"),0!==t.lastIndex||0!==e.lastIndex}(),c=n.UNSUPPORTED_Y||n.BROKEN_CARET,p=void 0!==/()??/.exec("")[1],f=l||p||c;f&&(i=function(t){var e,o,n,i,f=this,d=c&&f.sticky,u=r.call(f),h=f.source,v=0,g=t;return d&&(u=u.replace("y",""),-1===u.indexOf("g")&&(u+="g"),g=String(t).slice(f.lastIndex),f.lastIndex>0&&(!f.multiline||f.multiline&&"\n"!==t[f.lastIndex-1])&&(h="(?: "+h+")",g=" "+g,v++),o=new RegExp("^(?:"+h+")",u)),p&&(o=new RegExp("^"+h+"$(?!\\s)",u)),l&&(e=f.lastIndex),n=s.call(d?o:f,g),d?n?(n.input=n.input.slice(v),n[0]=n[0].slice(v),n.index=f.lastIndex,f.lastIndex+=n[0].length):f.lastIndex=0:l&&n&&(f.lastIndex=f.global?n.index+n[0].length:e),p&&n&&n.length>1&&a.call(n[0],o,(function(){for(i=1;i<arguments.length-2;i++)void 0===arguments[i]&&(n[i]=void 0)})),n}),t.exports=i},"9f7fd":function(t,e,o){"use strict";var r=o("d039");function n(t,e){return RegExp(t,e)}e.UNSUPPORTED_Y=r((function(){var t=n("a","y");return t.lastIndex=2,null!=t.exec("abcd")})),e.BROKEN_CARET=r((function(){var t=n("^r","gy");return t.lastIndex=2,null!=t.exec("str")}))},ac1f:function(t,e,o){"use strict";var r=o("23e7"),n=o("9263");r({target:"RegExp",proto:!0,forced:/./.exec!==n},{exec:n})},ad6d:function(t,e,o){"use strict";var r=o("825a");t.exports=function(){var t=r(this),e="";return t.global&&(e+="g"),t.ignoreCase&&(e+="i"),t.multiline&&(e+="m"),t.dotAll&&(e+="s"),t.unicode&&(e+="u"),t.sticky&&(e+="y"),e}},d784:function(t,e,o){"use strict";o("ac1f");var r=o("6eeb"),n=o("d039"),s=o("b622"),a=o("9263"),i=o("9112"),l=s("species"),c=!n((function(){var t=/./;return t.exec=function(){var t=[];return t.groups={a:"7"},t},"7"!=="".replace(t,"$<a>")})),p=function(){return"$0"==="a".replace(/./,"$0")}(),f=s("replace"),d=function(){return!!/./[f]&&""===/./[f]("a","$0")}(),u=!n((function(){var t=/(?:)/,e=t.exec;t.exec=function(){return e.apply(this,arguments)};var o="ab".split(t);return 2!==o.length||"a"!==o[0]||"b"!==o[1]}));t.exports=function(t,e,o,f){var h=s(t),v=!n((function(){var e={};return e[h]=function(){return 7},7!=""[t](e)})),g=v&&!n((function(){var e=!1,o=/a/;return"split"===t&&(o={},o.constructor={},o.constructor[l]=function(){return o},o.flags="",o[h]=/./[h]),o.exec=function(){return e=!0,null},o[h](""),!e}));if(!v||!g||"replace"===t&&(!c||!p||d)||"split"===t&&!u){var m=/./[h],y=o(h,""[t],(function(t,e,o,r,n){return e.exec===a?v&&!n?{done:!0,value:m.call(e,o,r)}:{done:!0,value:t.call(o,e,r)}:{done:!1}}),{REPLACE_KEEPS_$0:p,REGEXP_REPLACE_SUBSTITUTES_UNDEFINED_CAPTURE:d}),w=y[0],x=y[1];r(String.prototype,t,w),r(RegExp.prototype,h,2==e?function(t,e){return x.call(t,this,e)}:function(t){return x.call(t,this)})}f&&i(RegExp.prototype[h],"sham",!0)}},e01a:function(t,e,o){"use strict";var r=o("23e7"),n=o("83ab"),s=o("da84"),a=o("5135"),i=o("861d"),l=o("9bf2").f,c=o("e893"),p=s.Symbol;if(n&&"function"==typeof p&&(!("description"in p.prototype)||void 0!==p().description)){var f={},d=function(){var t=arguments.length<1||void 0===arguments[0]?void 0:String(arguments[0]),e=this instanceof d?new p(t):void 0===t?p():p(t);return""===t&&(f[e]=!0),e};c(d,p);var u=d.prototype=p.prototype;u.constructor=d;var h=u.toString,v="Symbol(test)"==String(p("test")),g=/^Symbol\((.*)\)[^)]+$/;l(u,"description",{configurable:!0,get:function(){var t=i(this)?this.valueOf():this,e=h.call(t);if(a(f,t))return"";var o=v?e.slice(7,-1):e.replace(g,"$1");return""===o?void 0:o}}),r({global:!0,forced:!0},{Symbol:d})}},f0c8:function(t,e,o){}}]);