package main
type (
    Point int
    TreeNode struct {
	    left, right *TreeNode
	    value *Comparable
    }
)
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  TypeDeclarationsImpl
    PsiElement(KEYWORD_TYPE)('type')
    PsiWhiteSpace(' ')
    PsiElement(()('(')
    PsiWhiteSpace('\n')
    PsiWhiteSpace('    ')
    TypeSpecImpl
      TypeNameDeclaration(Point)
        PsiElement(IDENTIFIER)('Point')
      PsiWhiteSpace(' ')
      TypeNameImpl
        LiteralIdentifierImpl
          PsiElement(IDENTIFIER)('int')
    PsiWhiteSpace('\n')
    PsiWhiteSpace('    ')
    TypeSpecImpl
      TypeNameDeclaration(TreeNode)
        PsiElement(IDENTIFIER)('TreeNode')
      PsiWhiteSpace(' ')
      TypeStructImpl
        PsiElement(KEYWORD_STRUCT)('struct')
        PsiWhiteSpace(' ')
        PsiElement({)('{')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('\t    ')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('left')
          PsiElement(,)(',')
          PsiWhiteSpace(' ')
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('right')
          PsiWhiteSpace(' ')
          TypePointerImpl
            PsiElement(*)('*')
            TypeNameImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('TreeNode')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('\t    ')
        TypeStructFieldImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('value')
          PsiWhiteSpace(' ')
          TypePointerImpl
            PsiElement(*)('*')
            TypeNameImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('Comparable')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('    ')
        PsiElement(})('}')
    PsiWhiteSpace('\n')
    PsiElement())(')')
