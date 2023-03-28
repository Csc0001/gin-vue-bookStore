import request from '@/utils/request';

// 用户注册
const register = ({ name, tel, password }) => request.post('/user/register', { name, tel, password });

// 获取用户信息
const info = () => request.get('/user/info');

export default {
  register,
  info,
};
