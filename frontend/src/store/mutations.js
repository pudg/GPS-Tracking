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
            hide: false
        }
    });
    state.devices = devs;
}

export function setTrackedDevices(state, devices) {
    state.trackedDevices = devices;
}

export function setUser(state, user) {
    state.user = user;
}

export function unsetUser(state) {
    state.user = null;
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