import { useRouter } from 'vue-router';
import axios from 'axios';

var auth = {
    getTokenHeader() {
        const token = localStorage.getItem('token');

        var config = {
            headers: { Authorization: `Bearer ${token}` }
        };

        return config;
    },
    checkToken(redirect = true, callback = null) {
        const url = new URL(window.location.href);
        const api = (url.port == "5173") ? "http://localhost:3001" : "/api/v1";

        const config = this.getTokenHeader();

        const router = useRouter();
        axios
            .get(api + "/auth",
                config
            )
            .then((response) => {
                if (response.status === 204) {
                    if (callback !== null) {
                        callback(false);
                    }
                }
                else {
                    if (callback !== null) {
                        callback(true);
                    }
                }
            })
            .catch((err) => {
                if (err || !err) {
                    if (redirect === true) {
                        if (callback !== null) {
                            callback(false);
                        }
                        localStorage.removeItem("token");
                        router.replace("/login");
                    }
                    else {
                        if (callback !== null) {
                            callback(false);
                        }
                    }
                }
            });
    }
}

export default auth;