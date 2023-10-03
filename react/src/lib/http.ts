import axios from "axios";

const instance = axios.create({
  baseURL: "http://127.0.0.1:8082/api",
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
});

instance.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      window.location.replace("/auth/register");
      localStorage.removeItem("token");
      return;
    }

    return Promise.reject(error);
  },
);

export const setAuthToken = (token: string): void => {
  localStorage.setItem("token", token);
  axios.defaults.headers.common["Authorization"] = "Bearer " + token;
};

let token = localStorage.getItem("token");
if (token) {
  setAuthToken(token);
}

export { instance };
