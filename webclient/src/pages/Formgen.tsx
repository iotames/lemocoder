import Generator from 'fr-generator';
import './index.less';
import FormsList from "@/pages/FormsList"

const { Provider, Sidebar, Canvas, Settings } = Generator;

const Demo = () => {
  return (
    <div>
        <div><FormsList /></div>
    <div className="fr-generator-playground" style={{ height: '80vh' }}>
      <Provider
        onChange={data => console.log('data:change', data)}
        onSchemaChange={schema => console.log('schema:change', schema)}
      >
        <div className="fr-generator-container">
          <Settings />
          <Canvas />
          <Sidebar fixedName />
        </div>
        </Provider>
    </div>
    
    </div>

  );
};

export default Demo;