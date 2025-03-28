import { toast } from 'react-toastify';

import { NOTIFICATION_TYPES } from '../const/notificationTypes';

const toastConfig: any = {
  position: 'top-right',
  autoClose: 5000,
  hideProgressBar: false,
  closeOnClick: true,
  pauseOnHover: true,
  draggable: true,
  progress: undefined,
  theme: 'colored',
};

export const openNotification = ({ type = NOTIFICATION_TYPES.ERROR, message }: any) => {
  switch (type) {
    case NOTIFICATION_TYPES.ERROR:
      toast.error(message, toastConfig);
      break;
    case NOTIFICATION_TYPES.INFO:
      toast.info(message, toastConfig);
      break;
    case NOTIFICATION_TYPES.SUCCESS:
      toast.success(message, toastConfig);
      break;
    case NOTIFICATION_TYPES.WARNING:
      toast.warn(message, toastConfig);
      break;
    default:
      toast(message, toastConfig);
  }
};
