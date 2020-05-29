import axios from 'axios';

axios.interceptors.request.use(function (config) {
  if(localStorage.getItem('userInfo')) config.headers.Authorization = JSON.parse(localStorage.getItem('userInfo')).access_token;
  return config;
});

export function postLogin(state) {
  return axios.post('http://localhost:5000/v1/auth/login', state).catch(function (error) {
    return error.response;
  });
}

export function getGcpClusters() {
  return axios.get('http://localhost:5001/v1/gcp/clusters').catch(function (error) {
    return error.response;
  });
}

export function setCluster(state) {
  return axios.post('http://localhost:5001/v1/'+state.type+'/cluster', state).catch(function (error) {
    return error.response;
  });
}

export function getSecrets() {
  return axios.get('http://localhost:5001/v1/gcp/secrets').catch(function (error) {
    return error.response;
  });
}

export function setSecret(state) {
  return axios.post('http://localhost:5001/v1/'+state.type+'/secret', state).catch(function (error) {
    return error.response;
  });
}


