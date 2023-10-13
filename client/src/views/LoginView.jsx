import { useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import useLogin from "../stores/useStore";
import Cookies from "js-cookie";
import { useNavigate } from "react-router-dom";

function LoginView() {
  const navigate = useNavigate();
  const { setLoggedIn } = useLogin();

  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });

  function handleLogin() {
    const sentData = JSON.stringify(formData);

    fetch(`http://localhost:8080/login`, {
      method: "POST",
      body: sentData,
    })
      .then((res) => {
        const auth = res.headers.get("Authorization");
        const ref = res.headers.get("Refresh-Token");

        Cookies.set("Auth", auth, { expires: 1 / 24 });
        Cookies.set("Ref", ref, { expires: 8 });

        return res.json();
      })
      .then((data) => {
        localStorage.setItem("username", data.username);
        localStorage.setItem("name", data.name);
        setLoggedIn(true);
        navigate("/posts/");
      });
  }

  const onChange = (e) => {
    setFormData((prevState) => ({
      ...prevState,
      [e.target.id]: e.target.value,
    }));
  };

  return (
    <>
      <Logo />
      <div className="flex flex-row">
        <Navbar />
        <div className="h-[600px]">
          <div className="flex flex-col gap-5 w-min">
            <p className="text-2xl font-semibold">Login</p>
            <form action="" className="flex flex-col gap-5 w-[236px]">
              <input
                onChange={onChange}
                type="text"
                id="username"
                className="bg-white h-[40px] px-2 focus:outline-none"
                placeholder="Username"
              />

              <input
                onChange={onChange}
                type="password"
                id="password"
                className="bg-white h-[40px] px-2 focus:outline-none"
                placeholder="Password"
              />
            </form>

            <button
              className="text-lg font-semibold bg-white w-[236px] h-[40px] hover:bg-[#FFA360]"
              onClick={handleLogin}
            >
              Login
            </button>
          </div>
        </div>
      </div>
    </>
  );
}

export default LoginView;
