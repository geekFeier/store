<template>
  <div id="app">
    <Menu mode="horizontal" :theme="theme1" active-name="1">
      <Row>
        <Col span="21">
          <MenuItem name="1">
            <a href="http://store.lameleg.com">
              <li>SealYun</li>
            </a>
          </MenuItem>
          <MenuItem name="2">
            <a target="_blank" href="https://sealyun.com">
              <li>Blog</li>
            </a>
          </MenuItem>
          <MenuItem name="3">
            <a target="_blank" href="https://sealyun.com/post/docs/">
              <li>Docs</li>
            </a>
          </MenuItem>
        </Col>
        <Col span="3">

          <MenuItem name="4">
            <Tooltip :content="'分享收入可提现金额,您可提现'+amount+'元'">
              <li @click="payeeFormCheck">
              <Modal
                v-model="payeeForm"
                @on-ok="ok"
                title="设置收款支付宝账号与提现密码"
                @on-cancel="cancel">
                <div class="payee">
                  <label>提现账号</label><input v-model="account" placeholder="收款支付宝账号"> <br/>
                  <label>提现密码</label><input v-model="passwd" placeholder="提现安全密码"><br/>
                  <label>确认密码</label><input v-model="passwdCheck" placeholder="切勿使用支付宝密码"><br/>
                </div>
               </Modal>
                <i class="ivu-icon ivu-icon-logo-yen">
                  <Badge :count="amount">&nbsp;&nbsp;&nbsp;&nbsp;</Badge>
                </i>
              </li>
            </Tooltip>
          </MenuItem>

          <MenuItem name="5">
            <img v-if="avata" :src="avata_url" style="border-radius:50%;width:30px;height:30px;cursor:pointer;margin-top:12px;"></img>
            <Tooltip v-else content="请使用github账户登录">
              <a
                href="https://github.com/login/oauth/authorize?client_id=89c1b05d77fb1c92a1ef&scope=user:email"
                target="_blank"
                rel="noopener"
              >
                <li>
                  <i class="ivu-icon ivu-icon-logo-github"></i>
                  Login
                </li>
              </a>
            </Tooltip>
          </MenuItem>
        </Col>
      </Row>
    </Menu>
    <HelloWorld msg="Welcome to SealYun"/>
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";

export default {
  data() {
    var a = {
      account:"",
      passwd:"",
      passwdCheck:"",

      amount: 1,
      theme1: "light",
      avata: false,
      avata_url: "",
      payeeForm: false,
    }
    this.$http.get('http://store.lameleg.com:8080/loginless/user/payee', { credentials: true } ).then(function(res){
              a.amount=res.data.Amount;
              console.log(res.data)
						},function(res){
              console.log(res.data)
            });

    this.$http.get('http://store.lameleg.com:8080/loginless/info/user', { credentials: true } ).then(function(res){
              a.avata_url=res.data.avatar_url;
//              a.avata_url="https://avatars2.githubusercontent.com/u/8912557?v=4"
              if (typeof(a.avata_url) != "undefined"){
                a.avata = true 
              }
              console.log(res.data,"avata:", a)
						},function(res){
              console.log(res.data)
            });

    return a;
  },
  name: "app",
  components: {
    HelloWorld
  },
  methods: {
    ok () {
      if (this.passwd != this.passwdCheck){
        this.$Message.info('两次密码设置不相同');
        return
      }

      var para = {
        PayeeAccount: this.account,
        Passwd: this.passwd,
      }

      this.$http.put("http://store.lameleg.com:8080/user/info/payee",para,{credentials:true,responseType:"json"}).then(function(res){
        console.log(res.data)
        this.$Message.info(res.data);
      },function(res){
        this.$Message.info(res.data);
      })
    },
    cancel () {
        this.$Message.info('取消设置');
    },
    payeeFormCheck:function(event){
      this.$http.post('http://store.lameleg.com:8080/user/info/withdraw',{},{credentials:true}).then(function(res){
        if (res.data.Amount == 0) {
          console.log("withdraw is 0")
        }
        if (res.data.PayeeAccount == "" || res.data.Passwd == ""){
          // if not set, else withdraw
          this.payeeForm = true
        }
        console.log(res.data)
      }, function(res){
        console.log("user info withdraw failed",res.data)
        this.payeeForm = true
      });
    },

    handleSubmit(name) {
        this.$refs[name].validate((valid) => {
            if (valid) {
                this.$Message.success('Success!');
            } else {
                this.$Message.error('Fail!');
            }
        })
    },
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.payee {
    margin: 10px;
}
.payee label {
  margin:10px;
}
.payee input{
  margin:10px;
}
</style>
