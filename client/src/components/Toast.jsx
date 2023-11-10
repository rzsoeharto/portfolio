import { toastStorage } from "../stores/useStore";
import Loader from "./atomic/Loader";

function Toast() {
  const { toastType, toastMessage } = toastStorage();

  let type;

  switch (toastType) {
    case "Loading":
      type = (
        <div className="flex flex-row gap-5 place-items-center place-self-center">
          <Loader />
          <p className="text-lg">{toastMessage}</p>
        </div>
      );
      break;

    case "Warning":
      type = (
        <>
          <div className="self-center">
            <p className="text-lg text-red-600">{toastMessage}</p>
          </div>
        </>
      );
      break;

    case "Success":
      type = (
        <>
          <div className="self-center">
            <p className="text-lg text-red-600">{toastMessage}</p>
          </div>
        </>
      );
      break;

    default:
      type = (
        <div className="">
          <p className="text-lg">{toastMessage}</p>
        </div>
      );
      break;
  }
  return (
    <>
      <div
        id="toast"
        className="flex bg-white absolute w-80 h-20 hidden font-semibold dropshadow-md bottom-[2%] left-[0%] px-5 place-content-center"
      >
        {type}
      </div>
    </>
  );
}

export default Toast;
