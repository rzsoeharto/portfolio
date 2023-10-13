import { useEffect, useState } from "react";
import Navbar from "../components/Navbar";
import TitleNav from "../components/atomic/TitleNav";
import formatPublishedDate from "../utils/dateFormatter";
import { Link } from "react-router-dom";
import Logo from "../components/Logo";
import Toast from "../components/toast";
import {
  closeToast,
  closeToastWithoutFade,
  showToast,
  toastWithoutFade,
} from "../utils/toastUtils";
import useLogin from "../stores/useStore";

function PostView() {
  const { isLoggedIn } = useLogin();
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    toastWithoutFade("Loading posts", "Loading");
    fetch("http://localhost:8080/posts")
      .then((response) => response.json())
      .then((data) => {
        setPosts(data);
        closeToast();
      })
      .catch((error) => {
        closeToastWithoutFade();
        console.log(error);
        showToast("Failed to fetch posts", "Warning");
      });
  }, []);

  return (
    <>
      <Logo />
      <div className="flex flex-row h-max w-full">
        <Navbar />
        <Toast />
        <div className="flex flex-col h-3/4 w-[540px]">
          <div
            className="flex flex-row w-full justify-between
          "
          >
            <TitleNav string="Entries" />
            {isLoggedIn ? (
              <Link to="/create-post" className="">
                <p className="text-3xl font-semibold w-[50px] text-center">+</p>
              </Link>
            ) : (
              <></>
            )}
          </div>
          <div className="flex gap-5 justify-between">
            <div className="flex flex-col gap-4">
              {Array.isArray(posts) &&
                posts.map((post, index) => (
                  <Link to={`/posts/${post.ID}`} key={index}>
                    <p className="text-xl duration-200 hover:text-[#FFA360] ">
                      {post.Title}
                    </p>
                  </Link>
                ))}
            </div>
            <div className="flex flex-col gap-4">
              {Array.isArray(posts) &&
                posts.map((post, index) => (
                  <p key={index} className="text-xl">
                    {formatPublishedDate(post.Published)}
                  </p>
                ))}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}

export default PostView;
