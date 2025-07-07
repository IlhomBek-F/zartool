import { Modal as AntModal } from 'antd';

type ModalProps = {
    children: React.ReactNode;
    handleConfirm: VoidFunction;
    handleClose: VoidFunction;
    isOpen: boolean;
}

function Modal({children, isOpen, handleClose, handleConfirm}: ModalProps) {
     
    return <AntModal
        open={isOpen}
        centered
        width={700}
        maskClosable={false}
        closable={false}
        okText="Сақлаш"
        onOk={handleConfirm}
        onCancel={handleClose}
        cancelText="Бекор қилиш"
      >
       {children}
      </AntModal>
}

export {Modal}