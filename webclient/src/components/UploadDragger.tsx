import {
    ProFormUploadDragger,
} from '@ant-design/pro-components';


  
const uploadField = (props:{onUploaded:(resp:any)=>void, name: string, action: string, label: string}) => {
    return (
        <>
            <ProFormUploadDragger max={4} label={props.label}  name="dragger" onChange={(info)=>{
            if (info.file.response != undefined){
                props.onUploaded(info.file.response)
            }
          }} />
        </>
    );
}

export default uploadField;