import { useEffect, useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import TitleNav from "../components/atomic/TitleNav";
import { useNavigate, useParams } from "react-router-dom";
import PostSection from "../components/PostSection";
import useLogin from "../stores/useStore";
import {
  closeToast,
  closeToastWithoutFade,
  showToast,
  toastWithoutFade,
} from "../utils/toastUtils";
import Cookies from "js-cookie";

function SpecificPostView() {
  const { isLoggedIn } = useLogin();
  const { postID } = useParams();
  const navigate = useNavigate();

  const [post, setPost] = useState([]);
  const [sections, setSections] = useState([]);

  function handleDelete() {
    const id = postID;
    const acc = Cookies.get("Auth");
    toastWithoutFade("Deleting post", "Loading");

    fetch(`http://localhost:8080/delete-post`, {
      method: "DELETE",
      headers: {
        Authorization: acc,
      },
      body: JSON.stringify({
        ID: Number(id),
      }),
    })
      .then((res) => {
        if (!res.ok) {
          closeToastWithoutFade();
          showToast("Failed to delete post", "Warning");
          return;
        }
        showToast("Post deleted", "Success");
        navigate("/posts");
      })
      .catch((error) => {
        closeToastWithoutFade();
        console.log(error);
        showToast("Unable to connect to server", "Warning");
      });
  }

  useEffect(() => {
    toastWithoutFade("", "Loading");
    fetch(`http://localhost:8080/post/${postID}`)
      .then((res) => res.json())
      .then((data) => {
        setPost(data);
        setSections(data.Sections);
      })
      .catch((error) => {
        closeToastWithoutFade();
        console.log(error);
        showToast("Failed to load post", "Warning");
      });
    closeToast();
  }, []);

  return (
    <>
      <Logo />
      <div className="flex flex-row h-min">
        <Navbar />
        <div className="flex flex-col w-[1080px]">
          <div className="flex flex-row justify-between h-min">
            <TitleNav string={post.Title} />
            {isLoggedIn ? (
              <button
                onClick={handleDelete}
                className="text-2xl bg-transparent duration-200 hover:text-red-600 font-semibold"
              >
                Delete post
              </button>
            ) : (
              <></>
            )}
          </div>
          <div className="flex flex-col gap-2">
            {Array.isArray(sections) &&
              sections.map((data, index) => (
                <div key={index} className="flex">
                  <PostSection sectionData={data} />
                </div>
              ))}
          </div>
        </div>
      </div>
    </>
  );
}

export default SpecificPostView;
