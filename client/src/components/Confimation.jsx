import { useNavigate, useParams } from "react-router-dom";
import { modalStorage } from "../stores/useStore";
import Cookies from "js-cookie";
import {
  closeToastWithoutFade,
  showToast,
  toastWithoutFade,
} from "../utils/toastUtils";

function Confimation() {
  const { postID } = useParams();
  const { modalType, setModalType, setModalState } = modalStorage();
  const navigate = useNavigate();

  function confirmLeave() {
    navigate("/posts");
    setModalState(false);
  }

  function handleClose() {
    setModalState(false);
  }

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
        setModalState(false);
        setModalType("");
        navigate("/posts");
      })
      .catch((error) => {
        closeToastWithoutFade();
        console.log(error);
        showToast("Unable to connect to server", "Warning");
      });
  }

  let type;

  switch (modalType) {
    case "Delete":
      type = (
        <>
          <div className="flex flex-col place-content-center font-semibold justify-around w-full h-full px-6">
            <p className="text-2xl">
              Are you sure you want to delete this post?
            </p>
            <div className="flex flex-row gap-5 place-content-end">
              <button
                className="w-[80px] h-[50px] bg-transparent text-blue-500 hover:text-blue-300 duration-300 "
                onClick={handleClose}
              >
                Cancel
              </button>
              <button
                className="w-[180px] h-[50px] bg-red-600 text-white rounded hover:text-black hover:bg-[#e7e7e7] duration-300"
                onClick={handleDelete}
              >
                Delete
              </button>
            </div>
          </div>
        </>
      );
      break;

    default:
      type = (
        <>
          <div className="flex flex-col place-content-center font-semibold justify-around w-full h-full px-6">
            <p className="text-2xl">
              You have unsaved changes <br />
              Are you sure you want to leave?
            </p>
            <div className="flex flex-row gap-5 place-content-end">
              <button
                className="w-[80px] h-[50px] hover:bg-red-500 duration-100"
                onClick={confirmLeave}
              >
                Leave
              </button>
              <button
                className="w-[180px] h-[50px] bg-[#ff8e3d] hover:bg-[#FFCCA8]"
                onClick={handleClose}
              >
                Stay
              </button>
            </div>
          </div>
        </>
      );
      break;
  }

  return (
    <>
      <div className="absolute z-50 bg-white left-[37%] top-[35%] w-[500px] h-[200px] rounded-md">
        {type}
      </div>
      <div
        onClick={handleClose}
        className="absolute z-30 bg-black w-full h-full opacity-20"
      ></div>
    </>
  );
}

export default Confimation;
