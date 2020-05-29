import axios from 'axios';

axios.interceptors.request.use(function (config) {
  if(localStorage.getItem('userInfo')) config.headers.Authorization = JSON.parse(localStorage.getItem('userInfo')).access_token;
  return config;
});

const cloud_url = 'http://localhost:5001/v1/cloud';
const auth_url = 'http://localhost:5000/v1/auth';

export function postLogin(state) {
  return axios.post(auth_url + '/login', state).catch(function (error) {
    return error.response;
  });
}

export function getGcpClusters() {
  return axios.get(cloud_url + '/clusters').catch(function (error) {
    return error.response;
  });
}

export function setCluster(state) {
  return axios.post(cloud_url + '/cluster', state).catch(function (error) {
    return error.response;
  });
}

export function getSecrets() {
  return axios.get(cloud_url + '/secrets').catch(function (error) {
    return error.response;
  });
}

export function setSecret(state) {
  return axios.post(cloud_url + '/secret', state).catch(function (error) {
    return error.response;
  });
}


