import axios from 'axios';

export function postLogin(id, pw) {
  return axios.post('http://localhost/login', {id:id, pw:pw});
}
