import { jwtDecode } from "jwt-decode";

export default function useAuth() {
  if (localStorage.getItem("token") === null || localStorage.getItem("token") === undefined) {
    return false;
  } else {
    try {
      jwtDecode(localStorage.getItem("token") ?? "");
      return true;
    } catch (e) {
      return false;
    }
  }
}
