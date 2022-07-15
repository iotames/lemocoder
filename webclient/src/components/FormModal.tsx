import { message, Modal } from 'antd';
import FormRender, { useForm } from 'form-render';
import type { FormInstance, Error, Schema as FormSchema } from 'form-render';
import { useState } from 'react';
import { post } from '@/services/api';

const AntdForm = (props:{form: FormInstance, formSchema: FormSchema, onFinish: (formData: any, error: Error[]) => void}) => {
    return (
        <FormRender form={props.form} schema={props.formSchema} onFinish={props.onFinish} />
    );
};

const FromModal = (props:{title: string, formSchema: FormSchema, postUrl: string, visible: boolean, setVisible: React.Dispatch<React.SetStateAction<boolean>>}) => {
  const visible = props.visible
  const setVisible = props.setVisible
  const [confirmLoading, setConfirmLoading] = useState(false);

  const handleSubmit = async (data: any, errors: Error[]) => {
    setConfirmLoading(true);
    console.log("----log----error", errors)
    if (errors.length > 0){
      message.error(
        '校验未通过：' + JSON.stringify(errors.map(item => {return item.error[0]}))
      )
      setConfirmLoading(false);
      return
    }
    console.log(data, errors)
    const result = await post(props.postUrl, data)
    setConfirmLoading(false);
    setVisible(false);
    console.log(result)
  }

  const handleCancel = () => {
    setVisible(false);
  };

  const form = useForm();
    return (
        <Modal
        title={props.title}
        visible={visible}
        onOk={form.submit}
        onCancel={handleCancel}
        confirmLoading={confirmLoading}
      >
        <AntdForm form={form} formSchema={props.formSchema} onFinish={handleSubmit} />
      </Modal>
    );
};

export default FromModal;

