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

export function setAuthenticated(state, authenticated) {

}

export function setTrack(status) {
    state.trackActive = status;
}