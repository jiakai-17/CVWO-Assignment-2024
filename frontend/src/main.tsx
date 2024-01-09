import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./index.css";
import App from "./App.tsx";
import MainPage from "./pages/MainPage.tsx";
import ThreadPage, { loader as ThreadLoader } from "./pages/ThreadPage.tsx";
import ThreadEditorPage from "./pages/ThreadEditorPage.tsx";
import AuthPage from "./pages/AuthPage.tsx";

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
        loader: ThreadLoader,
      },
      {
        path: "/viewthread/:id/edit",
        element: <ThreadEditorPage type={"edit"} />,
        loader: ThreadLoader,
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
        element: <h1>Not Found</h1>,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
);
