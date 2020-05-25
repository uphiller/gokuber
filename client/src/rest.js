import axios from 'axios';

/*axios.interceptors.request.use(function (config) {
  const token = store.getState().session.token;
  config.headers.Authorization =  token;
  console.log("aa");
  return config;
});*/

export function postLogin(state) {
  return axios.post('http://localhost:5000/v1/auth/login', state);
}
