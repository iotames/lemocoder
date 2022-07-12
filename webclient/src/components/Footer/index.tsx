import { DefaultFooter } from '@ant-design/pro-layout';

const Footer: React.FC = () => {
  const defaultMessage = 'AntDesign';
  const currentYear = new Date().getFullYear();

  return (
    <DefaultFooter
      copyright={`${currentYear} ${defaultMessage}`}
      links={[
        {
          key: 'key TODO',
          title: "TODO title",
          href: "#",
          blankTarget: true,
        },
      ]}
    />
  );
};

export default Footer;
