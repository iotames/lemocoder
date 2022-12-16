import {
    ProFormSwitch, 
    ProFormDigit, 
    ProFormText, 
    ProFormSelect, 
    ProFormGroup
} from '@ant-design/pro-components';


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
            <ProFormSelect name="ValueType" label="值的类型" tooltip="会生成不同的渲染器" initialValue="text" options={[
                {value:"text", label:"纯文本"}
                ]} rules={[{ required: true }]} />
            <ProFormText name="DataName" placeholder="字段名" label="字段名" rules={[{ required: true }]} width={120} tooltip="英文字母" /> 
            <ProFormText name="Title" placeholder="字段标题" label="字段标题" rules={[{ required: true }]} width={120} tooltip="中英文均可" />
            <ProFormDigit label="宽度(px)" min={0} max={300} name="Width" placeholder="像素宽度" width={90} tooltip="整数(默认0).范围:0~300" />
            <ProFormSwitch name="Editable" label="可编辑" />

            {/* 
            <ProFormSwitch name="Sorter" label="可排序" />
            <ProFormSwitch name="Copyable" label="可复制" />
            <ProFormSwitch name="Ellipsis" label="省略过长内容" />
            <ProFormDigit label="权重" tooltip="查询表单中的权重，权重大排序靠前" min={0} max={300} name="Order" placeholder={"请输入权重"} /> */}

      </ProFormGroup>
    );
} 