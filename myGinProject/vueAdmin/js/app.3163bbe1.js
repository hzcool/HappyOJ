(function(e){function t(t){for(var o,a,i=t[0],c=t[1],s=t[2],l=0,d=[];l<i.length;l++)a=i[l],Object.prototype.hasOwnProperty.call(r,a)&&r[a]&&d.push(r[a][0]),r[a]=0;for(o in c)Object.prototype.hasOwnProperty.call(c,o)&&(e[o]=c[o]);p&&p(t);while(d.length)d.shift()();return u.push.apply(u,s||[]),n()}function n(){for(var e,t=0;t<u.length;t++){for(var n=u[t],o=!0,a=1;a<n.length;a++){var i=n[a];0!==r[i]&&(o=!1)}o&&(u.splice(t--,1),e=c(c.s=n[0]))}return e}var o={},a={app:0},r={app:0},u=[];function i(e){return c.p+"js/"+({}[e]||e)+"."+{"chunk-00fdd667":"e43fa376","chunk-0c8a86fd":"063e8847","chunk-2b35aaf0":"a8ac5579","chunk-4d96cd5e":"9d7e50c7","chunk-52dfcaee":"37e1c976","chunk-74d83232":"9208e83e","chunk-7b3e8dac":"50e11f20","chunk-c62b9902":"ef446426"}[e]+".js"}function c(t){if(o[t])return o[t].exports;var n=o[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,c),n.l=!0,n.exports}c.e=function(e){var t=[],n={"chunk-2b35aaf0":1,"chunk-52dfcaee":1,"chunk-74d83232":1,"chunk-c62b9902":1};a[e]?t.push(a[e]):0!==a[e]&&n[e]&&t.push(a[e]=new Promise((function(t,n){for(var o="css/"+({}[e]||e)+"."+{"chunk-00fdd667":"31d6cfe0","chunk-0c8a86fd":"31d6cfe0","chunk-2b35aaf0":"858bb93a","chunk-4d96cd5e":"31d6cfe0","chunk-52dfcaee":"0a7d20de","chunk-74d83232":"858bb93a","chunk-7b3e8dac":"31d6cfe0","chunk-c62b9902":"2f2b2b7f"}[e]+".css",r=c.p+o,u=document.getElementsByTagName("link"),i=0;i<u.length;i++){var s=u[i],l=s.getAttribute("data-href")||s.getAttribute("href");if("stylesheet"===s.rel&&(l===o||l===r))return t()}var d=document.getElementsByTagName("style");for(i=0;i<d.length;i++){s=d[i],l=s.getAttribute("data-href");if(l===o||l===r)return t()}var p=document.createElement("link");p.rel="stylesheet",p.type="text/css",p.onload=t,p.onerror=function(t){var o=t&&t.target&&t.target.src||r,u=new Error("Loading CSS chunk "+e+" failed.\n("+o+")");u.code="CSS_CHUNK_LOAD_FAILED",u.request=o,delete a[e],p.parentNode.removeChild(p),n(u)},p.href=r;var f=document.getElementsByTagName("head")[0];f.appendChild(p)})).then((function(){a[e]=0})));var o=r[e];if(0!==o)if(o)t.push(o[2]);else{var u=new Promise((function(t,n){o=r[e]=[t,n]}));t.push(o[2]=u);var s,l=document.createElement("script");l.charset="utf-8",l.timeout=120,c.nc&&l.setAttribute("nonce",c.nc),l.src=i(e);var d=new Error;s=function(t){l.onerror=l.onload=null,clearTimeout(p);var n=r[e];if(0!==n){if(n){var o=t&&("load"===t.type?"missing":t.type),a=t&&t.target&&t.target.src;d.message="Loading chunk "+e+" failed.\n("+o+": "+a+")",d.name="ChunkLoadError",d.type=o,d.request=a,n[1](d)}r[e]=void 0}};var p=setTimeout((function(){s({type:"timeout",target:l})}),12e4);l.onerror=l.onload=s,document.head.appendChild(l)}return Promise.all(t)},c.m=e,c.c=o,c.d=function(e,t,n){c.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},c.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},c.t=function(e,t){if(1&t&&(e=c(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(c.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)c.d(n,o,function(t){return e[t]}.bind(null,o));return n},c.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return c.d(t,"a",t),t},c.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},c.p="/admin/",c.oe=function(e){throw console.error(e),e};var s=window["webpackJsonp"]=window["webpackJsonp"]||[],l=s.push.bind(s);s.push=t,s=s.slice();for(var d=0;d<s.length;d++)t(s[d]);var p=l;u.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"50fb":function(e,t,n){"use strict";var o=n("f684"),a=n.n(o);a.a},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var o=n("2b0e"),a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("router-view")],1)},r=[],u=n("2877"),i={},c=Object(u["a"])(i,a,r,!1,null,null,null),s=c.exports,l=(n("d3b7"),n("8c4f")),d=function(){var e=this,t=e.$createElement,n=e._self._c||t;return e.isLogin?n("div",{staticClass:"layout"},[n("Layout",[n("Sider",{ref:"side1",attrs:{"hide-trigger":"",collapsible:"","collapsed-width":78},model:{value:e.isCollapsed,callback:function(t){e.isCollapsed=t},expression:"isCollapsed"}},[n("Menu",{class:e.menuitemClasses,attrs:{"active-name":"1",theme:"dark",width:"auto"}},[n("MenuItem",{attrs:{name:"1"}},[n("Icon",{attrs:{type:"ios-home-outline"}}),n("span",[e._v("主页")])],1),n("MenuItem",{attrs:{name:"2"}},[n("Icon",{attrs:{type:"ios-people-outline"}}),n("span",[e._v("用户管理")])],1),n("MenuItem",{attrs:{name:"3",to:{name:"problemManage"}}},[n("Icon",{attrs:{type:"ios-help-circle-outline"}}),n("span",[e._v("题库管理")])],1),n("MenuItem",{attrs:{name:"4",to:{name:"contestManage"}}},[n("Icon",{attrs:{type:"ios-podium-outline"}}),n("span",[e._v("比赛管理")])],1)],1)],1),n("Layout",[n("Header",{staticClass:"layout-header-bar",style:{padding:0}},[n("Row",{style:{background:"#fff",boxShadow:"0 2px 3px 2px rgba(0,0,0,.1)",height:"80px"}},[n("Icon",{class:e.rotateIcon,style:{margin:"28px 20px",float:"left"},attrs:{type:"md-menu",size:"24"},nativeOn:{click:function(t){return e.collapsedSider(t)}}}),n("div",{staticStyle:{float:"left","margin-top":"5px"}},[n("span",{staticStyle:{"margin-left":"20px","font-size":"22px",color:"black"}},[e._v("HOJ后台管理系统")])]),n("div",{staticStyle:{float:"right","margin-top":"5px"}},[n("Button",{attrs:{type:"text"}},[e._v(e._s(e.name))]),n("Button",{staticStyle:{"margin-right":"15px"},attrs:{type:"error"},on:{click:e.logout}},[e._v("退出")])],1),n("div",{staticClass:"clear"})],1)],1),n("Content",{style:{margin:"20px",background:"#fff",minHeight:"1000px"}},[n("router-view"),n("br"),n("Footer",{staticClass:"layout-footer-center",staticStyle:{"text-align":"center"}},[e._v("author : hzcool "),n("br"),e._v(" "+e._s(e.date))])],1)],1)],1)],1):e._e()},p=[],f=(n("b0c0"),{components:{},data:function(){return{isCollapsed:!1,date:new Date}},computed:{isLogin:function(){return this.$store.state.isLogin},name:function(){return this.$store.state.name},rotateIcon:function(){return["menu-icon",this.isCollapsed?"rotate-icon":""]},menuitemClasses:function(){return["menu-item",this.isCollapsed?"collapsed-menu":""]}},methods:{collapsedSider:function(){this.$refs.side1.toggleCollapse()},logout:function(){var e=this;this.$req.get({url:"logout"}).then((function(t){e.$message("success",t),e.$turnToLogin(),e.$store.commit("Logout")}))}},mounted:function(){this.$store.state.isLogin||this.$req.autologin();var e=this;this.timer=setInterval((function(){e.date=new Date}),1e3)},beforeDestroy:function(){this.timer&&clearInterval(this.timer)}}),m=f,h=(n("50fb"),Object(u["a"])(m,d,p,!1,null,"170274fd",null)),g=h.exports;o["default"].use(l["a"]);var b=[{path:"/",name:"home",component:g,children:[{path:"problemManage",name:"problemManage",component:function(){return n.e("chunk-7b3e8dac").then(n.bind(null,"000a"))}},{path:"problemInfo",name:"problemInfo",component:function(){return n.e("chunk-74d83232").then(n.bind(null,"3424"))}},{path:"testdata",name:"testdata",component:function(){return n.e("chunk-0c8a86fd").then(n.bind(null,"394ca"))}},{path:"contestManage",name:"contestManage",component:function(){return n.e("chunk-52dfcaee").then(n.bind(null,"c4e7"))}},{path:"contestInfo",name:"contestInfo",component:function(){return n.e("chunk-2b35aaf0").then(n.bind(null,"4b41"))}},{path:"cproblems",name:"cproblems",component:function(){return n.e("chunk-4d96cd5e").then(n.bind(null,"0899"))}},{path:"csubmissions",name:"csubmissions",component:function(){return n.e("chunk-c62b9902").then(n.bind(null,"d8bc"))}}]},{path:"/login",name:"login",component:function(){return n.e("chunk-00fdd667").then(n.bind(null,"dd7b"))}}],v=new l["a"]({mode:"history",base:"/admin/",routes:b}),y=v,k=n("2f62");o["default"].use(k["a"]);var w=new k["a"].Store({state:{name:"傻逼",avatar:"",isLogin:!1,loaddingReq:!1},mutations:{Login:function(e,t){e.name=t,e.isLogin=!0},Logout:function(e){e.name="",e.isLogin=!1},Loadding:function(e){e.loaddingReq=!0},LoaddingCancle:function(e){e.loaddingReq=!1}},actions:{},modules:{}}),x=(n("96cf"),n("1da1")),C=n("bc3a"),I=n.n(C);I.a.defaults.headers.post["Content-Type"]="application/x-www-fromurlencodeed";var M=function(){y.push({name:"login"})},S={post:function(e){var t=e.url,n=e.form,o=e.config;return Object(x["a"])(regeneratorRuntime.mark((function e(){var a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,I.a.post(t,n,o).then((function(e){a=e.data})).catch((function(e){401==e.response.status&&y.push({name:"login"}),e.response.data&&alert(e.response.data)}));case 2:return e.abrupt("return",a);case 3:case"end":return e.stop()}}),e)})))()},get:function(e){var t=e.url,n=e.params;return Object(x["a"])(regeneratorRuntime.mark((function e(){var o;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,I.a.get(t,{params:n}).then((function(e){o=e.data})).catch((function(e){401==e.response.status&&y.push({name:"login"}),e.response.data&&alert(e.response.data)}));case 2:return e.abrupt("return",o);case 3:case"end":return e.stop()}}),e)})))()},autologin:function(){return Object(x["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,I.a.get("admin/autologin").then((function(e){w.commit("Login",e.data)})).catch((function(e){console.log(e.response.data),y.push({name:"login"})}));case 2:case"end":return e.stop()}}),e)})))()}},L=function(e){e.prototype.$req=S,e.prototype.$turnToLogin=M},O={install:L},_=n("a7fe"),P=n.n(_),T=(n("4160"),n("b64b"),n("159b"),n("5530")),j=(n("f8ce"),n("f825")),R={Affix:j["Affix"],Avatar:j["Avatar"],Badge:j["Badge"],Button:j["Button"],Card:j["Card"],Checkbox:j["Checkbox"],CheckboxGroup:j["CheckboxGroup"],Col:j["Col"],Content:j["Content"],Dropdown:j["Dropdown"],DropdownItem:j["DropdownItem"],DropdownMenu:j["DropdownMenu"],Divider:j["Divider"],Footer:j["Footer"],Form:j["Form"],FormItem:j["FormItem"],Header:j["Header"],Icon:j["Icon"],Input:j["Input"],InputNumber:j["InputNumber"],Sider:j["Sider"],Submenu:j["Submenu"],Layout:j["Layout"],List:j["List"],ListItem:j["ListItem"],Menu:j["Menu"],MenuGroup:j["MenuGroup"],MenuItem:j["MenuItem"],Message:j["Message"],Modal:j["Modal"],Option:j["Option"],Page:j["Page"],Panel:j["Panel"],Progress:j["Progress"],Radio:j["Radio"],RadioGroup:j["RadioGroup"],Row:j["Row"],Select:j["Select"],Spin:j["Spin"],Table:j["Table"],Tabs:j["Tabs"],TabPane:j["TabPane"],Upload:j["Upload"],DatePicker:j["DatePicker"]},D=Object(T["a"])({},R,{iButton:j["Button"],iCol:j["Col"],iContent:j["Content"],iForm:j["Form"],iFooter:j["Footer"],iHeader:j["Header"],iInput:j["Input"],iMenu:j["Menu"],iOption:j["Option"],iProgress:j["Progress"],iSelect:j["Select"],iTable:j["Table"]});Object.keys(D).forEach((function(e){o["default"].component(e,D[e])}));var $=function(e,t){j["Message"][e]({background:!0,content:t})};o["default"].prototype.$message=$;var E=n("b2d8"),F=n.n(E),B=(n("64e1"),n("8a2b")),A=n.n(B);n("1509");o["default"].use(P.a,I.a),o["default"].use(O),o["default"].use(F.a),o["default"].use(A.a),o["default"].config.productionTip=!1,new o["default"]({router:y,store:w,render:function(e){return e(s)}}).$mount("#app")},f684:function(e,t,n){}});