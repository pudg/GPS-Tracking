import axios from 'axios';
import router from '../router';

const axiosInstance = axios.create({
    withCredentials: true
});

export function saveUserPreferences({ commit }, preferences) {
    let devs = preferences.devices.map((dev) => {
        return JSON.stringify({
            id: dev.id,
            hide: dev.hide,
            image: dev.imagePreview,
        });
    });
    const prefs = {
        userID: preferences.user,
        sortAsc: preferences.sort,
        devices: devs,
    };

    const headers = {
        'Content-Type': 'application/json'
    };

    axiosInstance.put("http://localhost:8000/api/preferences", prefs, {headers: headers})
    .then((resp) => {})
    .catch(err => console.error(err))
    .finally(() => {})
}

export function userAuthenticate({ commit }, credentials) {
    const user = {
        email: credentials.email,
        password: credentials.password
    };

    const headers = {
        'Content-Type': 'application/json',
    };

    axiosInstance.post('http://localhost:8000/api/login', user, {headers: headers})
    .then((resp) => {
        localStorage.setItem("userID", resp.data.data);
        commit('setUser', resp.data.data);
        commit('setLoginError', "");
        router.push({name: 'tracking'});
    })
    .catch((err) => {
        commit('setLoginError', "Invalid Email or Password.");
    })
    .finally(() => {})
}

export function userLogout({ commit }, credentials) {
    const data = {
        userID: credentials
    }
    const headers = {
        'Content-Type': 'application/json',
    };

    axiosInstance.post('http://localhost:8000/api/logout', data, {headers: headers})
    .then((resp) => {
        localStorage.removeItem("user");
        localStorage.removeItem("userDevices");
        commit('unsetUser');
        commit('clearDevices');
    })
    .catch((err) => {})
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
    axios.post('http://localhost:8000/api/register', user, {headers: headers})
    .then((resp) => {
        if (resp.status === 201) {
            commit('setLoginError', "");
            commit('setRegistrationError', "");
            router.push({name: 'login'});
        }
    })
    .catch((err) => {
        commit('setRegistrationError', "Email already in use.");
    })
    .finally(() => {})
}

export function searchDevices({ commit }) {
    axiosInstance.get('http://localhost:8000/api/devices')
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

export function updateDeviceImage({ commit }, device) {
    commit('setDeviceImage', device);
}

export function deviceSort({ commit }) {
    commit('setDeviceSort')
}

export function persistUser({ commit }, user) {
    commit('setUser', user);
}
