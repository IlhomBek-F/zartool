import { notification } from "antd";

export function useNotification() {
    const [api, contextHolder] = notification.useNotification();
    
    const success = (description: string) => {
      api.success({
        message: "Success",
        description: description
      })
    }

    const error = (description: string) => {
      api.error({
        message: "Error",
        description: description
      })
    }

    return {contextHolder, success, error}
}