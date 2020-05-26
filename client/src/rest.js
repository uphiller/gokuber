import axios from 'axios';
import {NotificationContainer, NotificationManager} from 'react-notifications';

axios.interceptors.request.use(function (config) {
  if(localStorage.getItem('userInfo')) config.headers.Authorization = JSON.parse(localStorage.getItem('userInfo')).access_token;
  return config;
});

export function postLogin(state) {
  return axios.post('http://localhost:5000/v1/auth/login', state).catch(function (error) {
    return error.response;
  });
}
