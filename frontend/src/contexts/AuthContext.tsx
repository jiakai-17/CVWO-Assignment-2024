import { createContext } from "react";

const AuthContext = createContext({
  auth: {
    // The object that describes the current authentication status
    username: "",
    token: "",
    iat: 0,
    exp: 0,
    isLogin: false,
  },
  setAuthFromToken: (token: string) => {
    // Set the authentication status from a token
    ((x) => x)(token);
  },
  resetAuth: () => {
    // Reset the authentication status to empty
  },
  // Whether the context has fully loaded
  isLoaded: false,
});

export default AuthContext;
