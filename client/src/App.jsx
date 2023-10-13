import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import "./App.css";
import PostView from "./views/PostView";
import HomeView from "./views/Home";
import ContactView from "./views/ContactView";
import PortfolioView from "./views/PortfolioView";
import SpecificPostView from "./views/SpecificPostView";
import LoginView from "./views/LoginView";
import { useEffect } from "react";
import Cookies from "js-cookie";
import useLogin from "./stores/useStore";
import NewPostView from "./views/NewPostView";
import PrivateRoute from "./components/PrivateRoute";
import Toast from "./components/toast";
import {
  closeToast,
  closeToastWithoutFade,
  showToast,
  toastWithoutFade,
} from "./utils/toastUtils";

function App() {
  const { setLoggedIn } = useLogin();

  function replenishToken() {
    toastWithoutFade("Authenticating", "Loading");
    const refresh = Cookies.get("Ref");

    if (refresh == null || refresh === "null") {
      Cookies.remove("Auth");
      Cookies.remove("Ref");
      localStorage.clear();
      return;
    }

    fetch("http://localhost:8080/replenish", {
      method: "POST",
      headers: {
        Authorization: refresh,
      },
    })
      .then((res) => {
        if (!res.ok) {
          setLoggedIn(false);
          showToast("Error replenishing token", "Warning");
          return;
        }

        const auth = res.headers.get("Authorization");
        const ref = res.headers.get("Refresh-Token");

        Cookies.set("Auth", auth, { expires: 1 / 24 });
        Cookies.set("Ref", ref, { expires: 8 });

        setLoggedIn(true);
        closeToast(1000);
      })
      .catch((error) => {
        closeToastWithoutFade();
        error.log(error);
        showToast("Error Authenticating", "Warning");
      });
  }

  useEffect(() => {
    replenishToken();
    const intervalReplenish = setInterval(replenishToken, 3000000);
    return () => clearInterval(intervalReplenish);
  }, []);

  return (
    <>
      <Toast />
      <Router>
        <Routes>
          <Route path="/" element={<HomeView />} />
          <Route path="/posts" element={<PostView />} />
          <Route path="/posts/:postID" element={<SpecificPostView />} />
          <Route path="/create-post" element={<PrivateRoute />}>
            <Route path="/create-post" element={<NewPostView />} />
          </Route>
          <Route path="/contact" element={<ContactView />} />
          <Route path="/portfolio" element={<PortfolioView />} />
          <Route path="/login" element={<LoginView />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
