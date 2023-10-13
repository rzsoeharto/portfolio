import { useNavigate } from "react-router-dom";
import { modalStorage } from "../stores/useStore";

function Confimation() {
  const { modalType, setModalState } = modalStorage();
  const navigate = useNavigate();

  function confirmLeave() {
    navigate("/posts");
    setModalState(false);
  }

  function handleClose() {
    setModalState(false);
  }

  let type;

  switch (modalType) {
    case "Delete":
      type = (
        <>
          <div className="flex flex-col place-content-center font-semibold justify-around w-full h-full px-6">
            <p className="text-2xl">Delete this post?</p>
            <div className="flex flex-row gap-5 place-content-end">
              <button
                className="w-[180px] h-[50px] bg-red-500 hover:bg-white"
                onClick{}
              >
                Delete
              </button>
              <button
                className="w-[80px] h-[50px] bg-[#ff8e3d] hover:bg-[#FFCCA8]  duration-100"
                onClick={}
              >
                Cancel
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
