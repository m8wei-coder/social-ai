import React, { forwardRef } from "react";
import { Form, Upload, Input } from "antd";
import { InboxOutlined } from "@ant-design/icons";

export const PostForm = forwardRef((props, formRef) => {
  //TODO: mobile responsiveness
  const formItemLayout = {
    labelCol: { span: 6 },
    wrapperCol: { span: 14 },
  };

  const normFile = (e) => (Array.isArray(e) ? e : e?.fileList);

  return (
    <Form name="validate_other" {...formItemLayout} ref={formRef}>
      <Form.Item
        name="description"
        label="Message"
        rules={[
          {
            required: true,
            message: "Please input your message!",
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item label="Dragger">
        <Form.Item
          name="uploadPost"
          valuePropName="fileList"
          getValueFromEvent={normFile}
          noStyle
          rules={[
            {
              required: true,
              message: "Please select an image/video!",
            },
          ]}
        >
          {/* 图片/视频检查 */}
          <Upload.Dragger name="files" beforeUpload={() => false}>
            <p className="ant-upload-drag-icon">
              <InboxOutlined />
            </p>
            <p className="ant-upload-text">
              Click or drag file to this area to upload
            </p>
          </Upload.Dragger>
        </Form.Item>
      </Form.Item>
    </Form>
  );
});
