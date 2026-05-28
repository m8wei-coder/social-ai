import React from "react";
import { useNavigate } from "react-router-dom";
import { Form, Input, Button, message } from "antd";
import axios from "axios";

import { BASE_URL } from "../constants";

// (mobile) responsiveness design
const formItemLayout = {
  labelCol: {
    xs: { span: 24 },
    sm: { span: 8 },
  },
  wrapperCol: {
    xs: { span: 24 },
    sm: { span: 16 },
  },
};

const tailFormItemLayout = {
  wrapperCol: {
    xs: {
      span: 16,
      offset: 0,
    },
    sm: {
      span: 16,
      offset: 8,
    },
  },
};

function Register(props) {
    const [form] = Form.useForm();
    const navigate = useNavigate();

    const onFinish = (values) => {
        const { username, password } = values;
        const options = {
            method: "POST",
            url: `${BASE_URL}/signup`,
            data: {
                username,
                password,
            },
            headers: {
                "Content-Type": "application/json",
            },
        };

        axios(options)
            .then(res => {
                if(res.status === 200) {
                    message.success("Register succeeded");
                    navigate("/login");
                }
            })
            .catch((err) => {
                message.error("Register failed");
            });
    };

    return <Form
        {...formItemLayout}
        form={form}
        name="register"
        onFinish={onFinish}
        className="register"
    >
        <Form.Item
            name="username"
            label="Username"
            rules={[
                {
                    required: true,
                    message: "Please input your username",
                },
            ]}>
            <Input />
        </Form.Item>
        <Form.Item
            name="password"
            label="Password"
            rules={[
                {
                    required: true,
                    message: "Please input your password",
                },
            ]}
            hasFeedback>
            <Input.Password />
        </Form.Item>

        <Form.Item
            name="confirm"
            label="Confirm Password"
            hasFeedback
            rules={[
                {
                    required: true,
                    message: "Please confirm your password",
                },
                ({ getFieldValue }) => ({
                    validator(rule, value) {
                        if (!value || getFieldValue("password") === value) {
                            return Promise.resolve();
                        }
                        return Promise.reject(
                            "The two passwords do not match"
                        );
                    },
                }),
            ]}>
            <Input.Password />
        </Form.Item>

        <Form.Item {...tailFormItemLayout}>
            <Button type="primary" htmlType="submit" ClassName="register-btn" style={{ backgroundColor: "black" }}>
                Register
            </Button>
        </Form.Item>        
    </Form>
}

export default Register;