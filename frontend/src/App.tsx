import Navbar from "./components/Navbar.tsx";
import { Outlet, ScrollRestoration } from "react-router-dom";
import { useEffect, useMemo, useState } from "react";
import AuthContext from "./contexts/AuthContext.tsx";
import { jwtDecode } from "jwt-decode";
import JWTPayload from "./models/JwtPayload.tsx";

function App() {
  // AuthContext

  const emptyAuth = useMemo(() => {
    return { username: "", token: "", iat: 0, exp: 0, isLogin: false };
  }, []);

  const [isLoaded, setIsLoaded] = useState(false);

  const [auth, setAuth] = useState(emptyAuth);
  const resetAuth = useMemo(
    () => () => {
      setAuth(emptyAuth);
      localStorage.removeItem("token");
    },
    [emptyAuth],
  );

  const setAuthFromToken = useMemo(
    () => (token: string) => {
      try {
        const decodedToken: JWTPayload = jwtDecode(token);
        const newAuth = {
          username: decodedToken.username,
          token: token,
          iat: decodedToken.iat,
          exp: decodedToken.exp,
          isLogin: true,
        };
        setAuth(newAuth);
      } catch (e) {
        resetAuth();
      }
    },
    [resetAuth],
  );

  // Automatically log in if token is found in local storage
  useEffect(() => {
    setIsLoaded(false);
    if (localStorage.getItem("token") !== null && localStorage.getItem("token") !== undefined) {
      setAuthFromToken(localStorage.getItem("token") ?? "");
    }
    setIsLoaded(true);
  }, [setAuthFromToken]);

  const value = useMemo(
    () => ({
      auth,
      setAuthFromToken,
      resetAuth,
      isLoaded: isLoaded,
    }),
    [auth, isLoaded, resetAuth, setAuthFromToken],
  );

  return (
    <AuthContext.Provider value={value}>
      <Navbar />
      <Outlet />
      <ScrollRestoration />
    </AuthContext.Provider>
  );
}

export default App;
