import { useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";

function LoginView() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });

  function handleLogin() {
    const sentData = JSON.stringify(formData);

    fetch(`http://localhost:8080/login`, {
      method: "POST",
      body: sentData,
    }).then((res) =>
      res.json().then((data) => {
        console.log(data);
      })
    );
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
        <div className="w-full h-[600px]">
          <div className="flex flex-col gap-5 w-min">
            <p className="text-2xl font-semibold">Login</p>
            <form action="" className="flex flex-col gap-5 w-[236px]">
              <input
                onChange={onChange}
                type="text"
                id="username"
                className="bg-white h-[40px] px-2"
                placeholder="Username"
              />

              <input
                onChange={onChange}
                type="password"
                id="password"
                className="bg-white h-[40px] px-2"
                placeholder="Password"
              />
            </form>

            <button
              className="text-lg font-semibold bg-[#FFA360] w-[236px] h-[40px]"
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
