export function setDeviceList(state, devices) {
    let devs = devices.result_list.map((device) => {
        return {
            id: device.device_id,
            position: {
                lat: device.latest_accurate_device_point.lat,
                lng: device.latest_accurate_device_point.lng,
            },
            hide: false
        }
    });
    state.devices = devs;
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