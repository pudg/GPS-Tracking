export function setDeviceList(state, devices) {
    let devs = devices.result_list.map((device) => {
        return {
            id: device.device_id,
            make: device.make,
            model: device.model,
            active: device.active_state,
            position: {
                lat: device.latest_accurate_device_point.lat,
                lng: device.latest_accurate_device_point.lng,
            },
            hide: false,
            imagePreview: null,
        }
    });
    state.devices = devs;
    localStorage.setItem("userDevices", JSON.stringify(devs));
}

export function reloadDeviceList(state, devices) {
    state.devices = devices;
}

export function setTrackedDevices(state, devices) {
    state.trackedDevices = devices;
}

export function setUser(state, user) {
    state.user = user;
}

export function unsetUser(state) {
    state.user = null;
    localStorage.clear();
}

export function clearDevices(state) {
    state.devices = [];
    state.hiddenDevices = {};
    state.trackedDevices = [];
}

export function setTrack(status) {
    state.trackActive = status;
}

export function setRegistrationError(state, status) {
    state.registrationError = status;
}

export function setLoginError(state, status) {
    state.loginError = status;
}

export function setHiddenDevices(state, id) {
    if (id in state.hiddenDevices) {
        let target = state.hiddenDevices[id];
        delete state.hiddenDevices[id];
        target.hide = !target.hide;
        state.trackedDevices.push(target);
    } else {
        let target = {};
        let updatedDevices = state.trackedDevices.filter((dev) => {
            if (dev.id == id) {
                target = dev;
            } else {
                return dev;
            }
        });
        target.hide = !target.hide;
        state.trackedDevices = updatedDevices;
        state.hiddenDevices[id] = target;
    }
}

export function setDeviceImage(state, device) {
    for (let dev of state.devices) {
        if (dev.id == device.id) {
            dev.imagePreview = device.image;
            break;
        }
    }
    localStorage.setItem("userDevices", JSON.stringify(state.devices));
}

export function setDeviceSort(state) {
    let sortedDevs = [];
    if (state.sortAsc) {
        sortedDevs = state.devices.sort((d1, d2) => {
            if (d1.model < d2.model) return -1;
            if (d1.model > d2.model) return 1;
            return 0;
        });
    } else {
        sortedDevs = state.devices.sort((d1, d2) => {
            if (d1.model < d2.model) return 1;
            if (d1.model > d2.model) return -1;
            return 0;
        });
    }
    state.devices = sortedDevs;
    state.sortAsc = !state.sortAsc;
}