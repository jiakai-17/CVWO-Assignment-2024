import { createContext } from "react";

const AuthContext = createContext({
  auth: {
    username: "",
    token: "",
    iat: 0,
    exp: 0,
    isLogin: false,
  },
  setAuthFromToken: (token: string) => {
    // dummy function to prevent typescript warning about unused variables
    ((x) => x)(token);
  },
  resetAuth: () => {},
  isLoaded: false,
});

export default AuthContext;
