import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./index.css";
import App from "./App.tsx";
import MainPage from "./pages/MainPage.tsx";
import ThreadPage from "./pages/ThreadPage.tsx";
import ThreadEditorPage from "./pages/ThreadEditorPage.tsx";
import AuthPage from "./pages/AuthPage.tsx";
import NotFound from "./pages/NotFound.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "/",
        element: <MainPage />,
      },
      {
        path: "/viewthread/:id",
        element: <ThreadPage />,
      },
      {
        path: "/viewthread/:id/edit",
        element: <ThreadEditorPage type={"edit"} />,
      },
      {
        path: "/new",
        element: <ThreadEditorPage type={"create"} />,
      },
      {
        path: "/login",
        element: <AuthPage type={"login"} />,
      },
      {
        path: "/signup",
        element: <AuthPage type={"signup"} />,
      },
      {
        path: "*",
        element: <NotFound />,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);
