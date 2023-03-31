<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                placeholder="输入用户名(可选)"
                required
              ></b-form-input>
            </b-form-group>
            <b-form-group label="电话">
              <b-form-input
                v-model="$v.user.tel.$model"
                type="number"
                placeholder="输入电话"
                required
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('tel')">
                手机号格式错误
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.password.$model"
                type="password"
                placeholder="输入密码"
                required
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码必须大于6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button
                @click="register"
                variant="outline-primary"
                block
              >注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators';
import customValidtor from '@/helper/validtor';
import { mapActions } from 'vuex';

export default {
  name: 'userRegister',
  data() {
    return {
      user: {
        name: '',
        tel: '',
        password: '',
      },
    };
  },
  validations: {
    user: {
      name: {

      },
      tel: {
        required,
        tel: customValidtor.telValidtor,
      },
      password: {
        required,
        minLength: minLength(6),
      },
    },
  },
  methods: {
    ...mapActions('userModule', { userRegister: 'register' }),
    validateState(name) {
      // es6解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      this.userRegister(this.user).then(() => {
        // 跳转主页
        this.$router.replace({ name: 'home' });
      }).catch((err) => {
        console.log(err);
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });
      console.log('register');
    },
  },
};
</script>

<style>
</style>
