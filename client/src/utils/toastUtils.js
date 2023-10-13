import { toastStorage } from "../stores/useStore";

let toastTimeout;

export function showToast(message, type) {
  var id = document.getElementById("toast");
  id.className = id.className.replace("hidden", "block show ");
  toastStorage.setState({ toastType: type, toastMessage: message });

  if (toastTimeout) {
    clearTimeout(toastTimeout);
  }

  toastTimeout = setTimeout(() => {
    id.className = id.className.replace("block ", "hidden show");
  }, 3000);
}

export function toastWithoutFade(message, type) {
  var id = document.getElementById("toast");
  id.className = id.className.replace("hidden", "block show ");
  toastStorage.setState({ toastType: type, toastMessage: message });
}

export function closeToast(delay) {
  var id = document.getElementById("toast");
  setTimeout(() => {
    id.className = id.className.replace("block ", "hidden show");
  }, delay);
}

export function closeToastWithoutFade() {
  var id = document.getElementById("toast");
  id.className = id.className.replace("block ", "hidden show");
}
