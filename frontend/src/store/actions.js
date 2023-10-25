import axios from 'axios';

export function userAuthenticate({ commit }, credentials) {
    console.log("Authenticating: ", credentials);
    const user = {
        email: credentials.email,
        password: credentials.password
    };

    const headers = {
        'Content-Type': 'application/json'
    };

    axios.post('http://localhost:8000/login', user, {headers: headers})
    .then((resp) => {
        console.log("Login response: ", resp.data);
        commit('setUser', user);
    })
    .catch(err => console.error(err))
    .finally(() => {})
}

export function userLogout({ commit }) {
    commit('logout')
}

export function searchDevices({ commit }) {
    axios.get('http://localhost:8000/devices')
    .then((resp) => {
        // console.log("Received: ", JSON.parse(resp.data.data));
        commit('setDeviceList', JSON.parse(resp.data.data));
    })
    .catch(err => console.error(err))
    .finally(() => {})
}

export function trackDevices({ commit }) {
    commit('setTrack', true)
}
