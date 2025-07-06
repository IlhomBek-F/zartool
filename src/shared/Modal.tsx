import { Modal as AntModal } from 'antd';

type ModalProps = {
    children: React.ReactNode;
    handleConfirm: (value: any) => void;
    handleClose: () => void;
    isOpen: boolean;
}

function Modal({children, isOpen, handleClose, handleConfirm}: ModalProps) {
     
    return <AntModal
        style={{ top: 20 }}
        open={isOpen}
        width={700}
        maskClosable={false}
        closable={false}
        okText="Сақлаш"
        onOk={() => handleConfirm(false)}
        onCancel={handleClose}
        cancelText="Бекор қилиш"
      >
       {children}
      </AntModal>
}

export {Modal}