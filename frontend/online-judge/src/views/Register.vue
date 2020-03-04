<template>
    <div :style="backgroundDiv">
        
        <div class="block">

            <Form :model="info">
                <FormItem style="color:#fff" >
                    <h2>Happy Online Judge</h2>
                        <Divider size="small"/>
                </FormItem>
                <FormItem>
                    <Row>
                        <Col span="4" style="color:#fff;font-size:20px">
                            用户名
                        </Col>
                        <Col span="18" offset="1">
                            <Input type="text" v-model="info.username" placeholder="username" size="large" autofocus="true">
                                <Icon type="ios-person-outline" slot="prepend"></Icon>
                            </Input>
                        </Col>
                    </Row>
                </FormItem>
                <br>
                <FormItem>
                    <Row>
                        <Col span="4" style="color:#fff;font-size:20px">
                            密 码
                        </Col>
                        <Col span="18" offset="1">
                            <Input type="password" v-model="info.password" placeholder="Password" size="large" @on-enter="login()">
                                <Icon type="ios-lock-outline" slot="prepend"></Icon>
                            </Input>
                        </Col>
                    </Row>
                </FormItem>
                <br>
                <FormItem>
                    <Row>
                        <Col span="3" offset="8">
                            <Button type="primary" size="large" @click="login">登陆</Button>
                        </Col>
                        <col span="3">
                            <Button type="success" size="large">注册</Button>
                        </Col>
                    </Row>
                </FormItem>
            </Form>
        </div>
        
    </div>
</template>

<script>
export default {
    name: 'login',
    data() {
        return {
            backgroundDiv : {
                backgroundImage:'url(' + require('@/assets/1.jpg') + ')',
                backgroundRepeat:'no-repeat',
                backgroundSize:'100% 100%',
                height : '1100px',
                padding : '5%',
            },
            info : {
                username : "",
                password : "",
            },
        }
    },
    methods : {
        validate(username,password) {
            if (username.length == 0) return "用户名不能为空";
            else if (username.length > 20) return "用户名太长";
            else if (username.search(" ")!=-1) return "用户名不能包含空格";
            else if(password.length == 0) return "密码不能为空";
            else if(password.length > 16) return "密码太长" ;
            else if(password.search(" ")!=-1) return "密码不能包含空格";
            return ""; 
        },

        login() {
            let err = this.validate(this.info.username,this.info.password);
            if (err != "")   {
                this.$message("error",err)
                return;
            } 
            this.$req.post({url:"login",form:JSON.stringify(this.info),config:{ headers:{ 'Content-Type': ' application/json' }}})
            .then ( res => {
                if (res.status != 200) {
                    this.$message("error",res.data);
                } else {
                    this.$message("success",res.data)
                    this.$store.commit("Login",this.info.username);
                    this.$router.push({name:"home"});
                } 
            });
        }
    }
}
</script>

<style scoped>
    .block {
        margin:0 auto;
        height:400px;
        background: rgba(0, 0, 0, 0.5);
        width: 500px;
        padding: 30px 60px;
    }
</style>