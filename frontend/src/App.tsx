import Navbar from "./components/Navbar.tsx";
import { Outlet, ScrollRestoration } from "react-router-dom";

function App() {
  return (
    <>
      <Navbar />
      <Outlet />
      <ScrollRestoration />
    </>
  );
}

export default App;
