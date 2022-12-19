import {
    ProFormSwitch, 
    ProFormDigit, 
    ProFormText, 
    ProFormSelect, 
    ProFormGroup
} from '@ant-design/pro-components';
import { useState } from 'react';


export const TableItemFormFields = () => {
    return (
        <ProFormGroup>
            <ProFormSelect name="DataType" label="字段类型" initialValue="string" options={[
                {value:"string", label:"字符串(string)"},
                {value:"int", label:"整型(int)"},
                {value:"float", label:"浮点型(float)"},
                {value:"text", label:"文本(text)"},
                {value:"smallint", label:"小整型(smallint)"},
                {value:"bigint", label:"长整型(bigint)"},
                ]} rules={[{ required: true }]} />
            <ProFormSelect name="ValueType" label="值的类型" tooltip="会生成不同的渲染器" hidden initialValue="text" options={[
                {value:"text", label:"纯文本"}
                ]} rules={[{ required: true }]} />
            <ProFormText name="DataName" placeholder="英文字段名" label="字段名" rules={[{ required: true }]} width={120} tooltip="英文字母" /> 
            <ProFormText name="Title" placeholder="中英文均可" label="字段标题" rules={[{ required: true }]} width={120} tooltip="中英文均可" />
            <ProFormDigit label="宽度(px)" min={0} max={300} name="Width" placeholder="像素宽度" width={90} tooltip="整数(默认0).范围:0~300" />
            <ProFormSwitch name="Editable" label="可编辑" initialValue={true} />

            {/* 
            <ProFormSwitch name="Sorter" label="可排序" />
            <ProFormSwitch name="Copyable" label="可复制" />
            <ProFormSwitch name="Ellipsis" label="省略过长内容" />
            <ProFormDigit label="权重" tooltip="查询表单中的权重，权重大排序靠前" min={0} max={300} name="Order" placeholder={"请输入权重"} /> */}

        </ProFormGroup>
    );
}

export const TableItemOptFormFields = () => {
    const [urlFieldHide, setUrlFieldHide] = useState(false);
    return (
        <ProFormGroup>
            <ProFormSelect name="Type" label="操作类型" initialValue="action"
                onChange={(value)=>{
                    if (value=="form" || value=="edit"){
                        setUrlFieldHide(true)
                    }
                    if (value=="action" || value=="redirect"){
                        setUrlFieldHide(false)
                    }
                }} 
            options={[
                {value:"action", label:"POST数据"},
                {value:"redirect", label:"路由跳转"},
                {value:"form", label:"表单提交"},
                {value:"edit", label:"快捷编辑"},
                ]} rules={[{ required: true }]} />
            <ProFormText name="Title" label="操作标题" rules={[{ required: true }]} width={90} />
            <ProFormText name="Key" label="操作名" rules={[{ required: true }]} width={90} tooltip="英文.更新行数据的表单提交, 操作名必须为update" />
            <ProFormText name="Url" label="地址" hidden={urlFieldHide} />
        </ProFormGroup>
    )
}