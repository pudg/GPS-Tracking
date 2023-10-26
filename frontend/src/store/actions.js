import axios from 'axios';
import router from '../router';

export function userAuthenticate({ commit }, credentials) {
    const user = {
        email: credentials.email,
        password: credentials.password
    };

    const headers = {
        'Content-Type': 'application/json'
    };

    axios.post('http://localhost:8000/login', user, {headers: headers})
    .then((resp) => {
        if (resp.status === 200) {
            commit('setUser', user);
            commit('setLoginError', "");
            router.push({name: 'tracking'});
        }
    })
    .catch((err) => {
        commit('setLoginError', "Invalid Email or Password.");
        console.error(err);
    })
    .finally(() => {})
}

export function userRegistration({ commit }, credentials) {
    const user = {
        email: credentials.email,
        password: credentials.password
    };
    const headers = {
        'Content-Type': 'application/json'
    }
    axios.post('http://localhost:8000/register', user, {headers: headers})
    .then((resp) => {
        if (resp.status === 201) {
            commit('setUser', user);
            commit('setRegistrationError', "");
            router.push({name: 'tracking'});
        }
    })
    .catch((err) => {
        commit('setRegistrationError', "Email already in use.");
    })
    .finally(() => {})
}

export function userLogout({ commit }) {
    commit('unsetUser')
}

export function searchDevices({ commit }) {
    axios.get('http://localhost:8000/devices')
    .then((resp) => {
        commit('setDeviceList', JSON.parse(resp.data.data));
    })
    .catch(err => console.error(err))
    .finally(() => {})
}

export function trackDevices({ commit }) {
    commit('setTrack', true)
}

export function hideDevice({ commit }, id) {
    commit('setHiddenDevices', id);
}
