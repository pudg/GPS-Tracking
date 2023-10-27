import axios from 'axios';
import router from '../router';
import store from '../store';


export function saveUserPreferences({ commit }, preferences) {
    let devs = preferences.devices.map((dev) => {
        return JSON.stringify({
            id: dev.id,
            hide: dev.hide,
            image: dev.imagePreview,
        });
    });
    const data = {
        email: preferences.user.email,
        password: preferences.user.password,
        preference: {
            sortAsc: preferences.sort,
            devices: devs,
        }
    };


    const headers = {
        'Content-Type': 'application/json'
    };

    axios.put("http://localhost:8000/api/preferences", data, {headers: headers})
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
        'Content-Type': 'application/json'
    };

    axios.post('http://localhost:8000/api/login', user, {headers: headers})
    .then((resp) => {
        console.log("statusCode: ", resp.status);
        commit('setUser', user);
        commit('setLoginError', "");
        router.push({name: 'tracking'});

    })
    .catch((err) => {
        commit('setLoginError', "Invalid Email or Password.");
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
    axios.post('http://localhost:8000/api/register', user, {headers: headers})
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

export function searchDevices({ commit }) {
    axios.get('http://localhost:8000/api/devices')
    .then((resp) => {
        commit('setDeviceList', JSON.parse(resp.data.data));
    })
    .catch(err => console.error(err))
    .finally(() => {})
}

export function userLogout({ commit }) {
    commit('unsetUser')
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