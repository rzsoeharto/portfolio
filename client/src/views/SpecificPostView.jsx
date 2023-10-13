import { useEffect, useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import TitleNav from "../components/atomic/TitleNav";
import { useParams } from "react-router-dom";
import PostSection from "../components/PostSection";
import useLogin from "../stores/useStore";
import {
  closeToast,
  closeToastWithoutFade,
  showToast,
  toastWithoutFade,
} from "../utils/toastUtils";

function SpecificPostView() {
  const { isLoggedIn } = useLogin();
  const { postID } = useParams();
  const [post, setPost] = useState([]);
  const [sections, setSections] = useState([]);

  function handleDelete() {
    console.log("ahjsbdhasbdhabsd");
  }

  useEffect(() => {
    toastWithoutFade("", "Loading");
    fetch(`http://localhost:8080/post/${postID}`)
      .then((response) => response.json())
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
